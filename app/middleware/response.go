package middleware

import (
	"github.com/gin-gonic/gin"
)

func ReponseAdmin()  gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Next()
		params := context.Keys
		response := params["response"].(gin.H)
		message := response["message"].(gin.H)
		//message["Username"] = ParseHeadOrCookie(context, utils.ASYUSERID)
		//message["Rolename"] = ParseHeadOrCookie(context, utils.ASTROLE)
		//message["Username"] = "Username"
		//message["Rolename"] = "管理员"
		context.HTML(response["code"].(int), response["template"].(string), message)
	}
}