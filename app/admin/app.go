package main

import (
	"encoding/gob"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"mxcms/app/models"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"html/template"
	"mxcms/app/admin/controller"
	"mxcms/mxcore/discovery"
	mxcoreutils "mxcms/mxcore/utils"

	"mxcms/app/middleware"
)

func SiteConfig() gin.H{
	return gin.H{
		"STATIC_URL":"public/",
		"SITE_URL":"public/admin/mx_admin",
		"MXCMS_SOFTNAME":"MXCMS",
		"MXCMS_VERSION":"0..0.1",
	}
}

func Start() {
	app := gin.New()
	app.Use(gin.Recovery())
	app.Use(middleware.TraceGin)

	initEtcd()

	initRouter(app)

	app.Run(":8080")
}

func main(){
	Start()
}


func initEtcd() {
	urls := discovery.GetEtcdUrls()
	reg := etcd.NewRegistry(func(op *registry.Options) {
		op.Addrs = urls
	})

	service := micro.NewService(
		micro.Registry(reg),
		micro.Name("micro.qgn.admin"),
	)
	//server.Init()
	//service.Server().Init(server.Wait(nil))
	go service.Run()
}


func initRouter(app *gin.Engine) {
	gob.Register(models.AdminUser{})
	store := cookie.NewStore([]byte("secret"))
	app.Use(sessions.Sessions("mxsession", store))

	app.Static("favicon.ico","./public/favicon.ico")

	initGroupLogin(app)
	initGroupIndex(app)

	app.SetFuncMap(template.FuncMap{
		"json2stringslice":mxcoreutils.Json2StringSlice,
		"SiteConfig":SiteConfig,
		"urlfor":mxcoreutils.URLFor(app.Routes()),
		"slice_int_get_key":mxcoreutils.SliceIntGetKey,
	})
	app.LoadHTMLGlob("app/admin/views/**/*")
}

func initGroupLogin(app *gin.Engine){
	app.Static("public","./public")

	app.GET("/login", controller.Login)
	app.POST("/login", controller.Login)
}

func initGroupIndex(app *gin.Engine)  {
	admin := app.Group("/index")
	admin.Static("public","./public")

	admin.Use(middleware.CheckLogin)
	admin.Use(middleware.ReponseAdmin())
	admin.GET("", controller.Index)
	admin.GET("/publichome", controller.PublicHome)
}
