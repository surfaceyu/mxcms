package admin

import (
	//"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"html/template"
	"mxcms/app/middleware"
	"mxcms/app/admin/controller"
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
func Start() {
	app := gin.New()
	app.Use(gin.Recovery())
	app.Use(middleware.TraceGin)

	app.Static("public","./public")
	app.Static("favicon.ico","./public/favicon.ico")
	app.SetFuncMap(template.FuncMap{
		"json2stringslice":mxcoreutils.Json2StringSlice,
		"SiteConfig":SiteConfig,
	})
	app.LoadHTMLGlob("app/admin/views/**/*")

	app.GET("/login", controller.Login)
	app.POST("/login", controller.Login)

	admin := app.Group("/")
	admin.Use(middleware.CheckLogin)
	admin.Use(middleware.Admininfo())
	admin.GET("/index", controller.Index)

	app.Run(":8080")
}

