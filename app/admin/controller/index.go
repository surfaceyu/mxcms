package controller

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"mxcms/app/models"
	"mxcms/app/rpccli"
	"net/http"
)

func Index(context *gin.Context) {
	res := new([]models.Msgmenu)
	_ = rpccli.Call("GetAdminMenu", 1, res)

	session := sessions.Default(context)
	adminUser:= session.Get("Admin")

	context.Set("response", gin.H{
		"code":http.StatusOK,
		"template":"admin/index/index.tmpl",
		"message":gin.H{
			"Menu" : res,
			"adminUser":adminUser,
		},
	})
}

func PublicHome(context *gin.Context) {
	//res := new([]models.Msgmenu)
	//_ = rpccli.Call("GetAdminMenu", 1, res)
	count := []int{
		2,4,6,8,
	}

	session := sessions.Default(context)
	adminUser:= session.Get("Admin")

	context.Set("response", gin.H{
		"code":http.StatusOK,
		"template":"admin/layouts/publichome",
		"message":gin.H{
			"count" : count,
			"admininfo":adminUser,
		},
	})
}