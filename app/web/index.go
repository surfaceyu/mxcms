package web

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"mxcms/app/middleware"
	"mxcms/app/web/controller"
	mxcoreutils "mxcms/mxcore/utils"
)

var app *gin.Engine

func SiteConfig() gin.H{
	return gin.H{
		"STATIC_URL":"",
		"SITE_URL":"",
	}
}

func Start(){
	app = gin.New()
	app.Use(gin.Recovery())
	app.Use(middleware.TraceGin)
	app.Static("public","./public")
	app.Static("favicon.ico","./public/favicon.ico")
	app.SetFuncMap(template.FuncMap{
		"json2stringslice":mxcoreutils.Json2StringSlice,
		"SiteConfig":SiteConfig,
	})
	app.LoadHTMLGlob("app/web/views/**/*")
	app.GET("/", controller.Index)
	//app.GET("/links/{id}", controller.ListLinks)
	//app.GET("/cuslinks", controller.ListCusLinks)
	//app.GET("/articles", controller.ListArticles)
	//app.GET("/article/{id}", controller.Detail)
	//app.GET("/payme", controller.Payme)
	//app.GET("/link/{id}", controller.GetLinkUrl)
	//app.GET("/cuslink/{id}", controller.GetCusLinkUrl)


	//api := app.Group("/")
	//api.Use(middleware.CheckLogin)
	//api.POST("/upload", controller.Upload)

	//article := api.Group("/article")
	//{
	//	article.POST("/save", controller.SaveArticle)
	//	article.POST("/list/:size/:page", controller.GetArticleList)
	//	article.GET("/get/:id", controller.GetArticle)
	//	article.GET("/delete/:id", controller.DeleteArticle)
	//}
	//
	//account := api.Group("/account")
	//{
	//	account.POST("/list/:size/:page", controller.GetAccountList)
	//	account.GET("/get/:id", controller.GetAccount)
	//}
	//
	//link := api.Group("/link")
	//{
	//	link.POST("/save", controller.SaveLink)
	//	link.POST("/list/:size/:page", controller.GetLinkList)
	//	link.GET("/get/:id", controller.GetLink)
	//	link.GET("/cat/list", controller.GetCatOptions)
	//	link.GET("/delete/:id", controller.DeleteLink)
	//}
	//
	//cuslink := api.Group("/cuslink")
	//{
	//	cuslink.POST("/save", controller.SaveCusLink)
	//	cuslink.POST("/list/:size/:page", controller.GetCusLinkList)
	//	cuslink.GET("/get/:id", controller.GetCusLink)
	//	cuslink.GET("/delete/:id", controller.DeleteCusLink)
	//}

	app.Run(":8090")
}

