package controller

import (
	"github.com/gin-gonic/gin"
	"mxcms/app/models"
	"mxcms/app/rpccli"
	"net/http"
)

func Index(context *gin.Context) {
	//platform, _ := json.Marshal([]string{"PS4","Xbox One","PC"})
	//good, _ := json.Marshal([]string{"A fitting extension of Dark Souls III","Impressive final boss","Diverse environment layouts","Lots of new enemies and gear to discover"})
	//bad, _ := json.Marshal([]string{"Doesn't leave a lasting impression","Abundant bonfires lessen exploration-based challenges"})
	//GameReview := dbmodels.GameReview{
	//	GameId: "da132das131",
	//	Name:"黑暗之魂3",
	//	Avatar:"https://gamespot1.cbsistatic.com/uploads/scale_avatar/1197/11970954/2976798-sss91mcjvyj85l._sl1500_.jpg",
	//	Platform:string(platform),
	//	Good:string(good),
	//	Bad:string(bad),
	//	Score:8.5,
	//	Scoreword:"好",
	//}
	//// 创建
	////app.Db.CreateTable(&dbmodels.GameReview{})
	//app.Db.Create(&GameReview)
	res := new([]models.Msgmenu)
	_ = rpccli.Call("GetAdminMenu", 1, res)
	context.Set("response", gin.H{
		"code":http.StatusOK,
		"template":"admin/index/index.tmpl",
		"message":gin.H{
			"Menu" : "res",
		},
	})
}