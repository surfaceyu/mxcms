package controller

import (
	"github.com/gin-gonic/gin"
	"mxcms/app/models"
	"mxcms/app/rpccli"
	"net/http"
)

func Index(context *gin.Context) {
	res := new([]models.Msgmenu)
	_ = rpccli.Call("GetAdminMenu", 1, res)
	context.Set("response", gin.H{
		"code":http.StatusOK,
		"template":"admin/index/index.tmpl",
		"message":gin.H{
			"Menu" : res,
		},
	})
}

func PublicHome(context *gin.Context) {
	//res := new([]models.Msgmenu)
	//_ = rpccli.Call("GetAdminMenu", 1, res)
	count := []int{
		2,4,6,8,
	}

	admininfo := models.AdminUser{
		1,"yzmcms","2894f6dbfdad0f278b5d2f339045a5f1",
		1,"超级管理员","","","",0,
		"0","1587363988","系统",
	}

	context.Set("response", gin.H{
		"code":http.StatusOK,
		"template":"admin/layouts/publichome",
		"message":gin.H{
			"count" : count,
			"admininfo":admininfo,
		},
	})
}