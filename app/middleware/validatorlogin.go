package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mxcms/mxcore/utils"
)

func CheckLogin(ctx *gin.Context) {
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
	ctx.Next()
}

func ParseHeadOrCookie(ctx *gin.Context, k string) string {
	v := ctx.GetHeader(k)
	if len(v) == 0 {
		v, _ = ctx.Cookie(k)
	}
	return v
}

func Admininfo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ParseHeadOrCookie(ctx, utils.ASYUSERID)
		//ctx.HTML(http.StatusOK, gin.H{
		//	Username:id,
		//})
		fmt.Println(ctx.Params, id)
		ctx.Next()
	}
}