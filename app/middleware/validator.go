package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"mxcms/app/models"
	"net/http"
)

func CheckLogin(context *gin.Context) {
	session := sessions.Default(context)
	adminUser:= session.Get("Admin")
	if adminUser != nil && adminUser.(models.AdminUser).Adminid > 0 {
		context.Next()
		return
	}
	context.Redirect(http.StatusMovedPermanently, "/login")
	//id := ParseHeadOrCookie(ctx, utils.ASYUSERID)
	//token := ParseHeadOrCookie(ctx, utils.ASYTOKEN)
	//
	//if len(id) <= 0 || len(token) <= 0 {
	//	log.Printf("miss id or token")
	//	return
	//}
	//
	//v, err := redis.Get(id)
	//if err != nil {
	//	log.Printf(errors.ErrorStack(err))
	//	return
	//}
	//
	//if v == token {
	//	ctx.Next()
	//}
	//context.Next()
}

func ParseHeadOrCookie(ctx *gin.Context, k string) string {
	v := ctx.GetHeader(k)
	if len(v) == 0 {
		v, _ = ctx.Cookie(k)
	}
	return v
}