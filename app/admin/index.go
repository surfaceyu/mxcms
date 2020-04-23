package admin

import (
	"github.com/gin-gonic/gin"
	"html/template"
	//"mxcms/app"
	"mxcms/app/admin/controller"
	"mxcms/app/middleware"
	mxcoreutils "mxcms/mxcore/utils"
)

func SiteConfig() gin.H{
	return gin.H{
		"STATIC_URL":"public/",
		"SITE_URL":"public/admin/mx_admin",
		"MXCMS_SOFTNAME":"MXCMS",
		"MXCMS_VERSION":"0..0.1",
	}
}

//func initEtcd() {
//	urls := discovery.GetEtcdUrls()
//	reg := etcd.NewRegistry(func(op *registry.Options) {
//		op.Addrs = urls
//	})
//
//	service := micro.NewService(
//		micro.Registry(reg),
//		micro.Name("micro.qgn.admin"),
//	)
//	//server.Init()
//	//service.Server().Init(server.Wait(nil))
//	go service.Run()
//}

func Start() {
	app := gin.New()
	app.Use(gin.Recovery())
	app.Use(middleware.TraceGin)

	app.Static("public","./public")
	app.Static("favicon.ico","./public/favicon.ico")

	app.GET("/login", controller.Login)
	app.POST("/login", controller.Login)

	admin := app.Group("/")
	admin.Use(middleware.CheckLogin)
	admin.Use(middleware.ReponseAdmin())
	admin.GET("/index", controller.Index)

	app.SetFuncMap(template.FuncMap{
		"json2stringslice":mxcoreutils.Json2StringSlice,
		"SiteConfig":SiteConfig,
		"urlfor":mxcoreutils.URLFor(app.Routes()),
	})
	app.LoadHTMLGlob("app/admin/views/**/*")
	app.Run(":8080")
}

func main(){
	Start()
}

