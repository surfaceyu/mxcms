package controller

import (
	"github.com/gin-gonic/gin"
	"mxcms/app/models"
	"mxcms/app/rpccli"
	"net/http"
)


func Index(context *gin.Context) {
	res := new(models.GameReview)
	_ = rpccli.Call("GetGameReview", "da132das131", res)
	//res := models.GameReview{
	//	GameId:"da132das131",
	//}
	//app.Db.First(&GameReview)
	context.HTML(http.StatusOK, "defaults/index.tmpl", gin.H{
		"title":"QGN-"+res.Name,
		"scorereview":res,
		"article":"",
	} )
}