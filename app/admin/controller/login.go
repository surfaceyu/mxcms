package controller

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"mxcms/app/models"
	"mxcms/app/rpccli"
	"mxcms/mxcore/mxconst"
	"mxcms/mxcore/utils"
	"net/http"
)

func Login(context *gin.Context) {
	if context.Request.Method == mxconst.POST {
		var form models.MsgLogin
		// This will infer what binder to use depending on the content-type header.
		if err := context.ShouldBind(&form); err != nil {
			context.HTML(http.StatusOK, "index/login.tmpl", gin.H{
				"MXCMS_SOFTNAME":"MXCMS",
				"Code":"1001",
			})
			return
		}
		res := new(models.AdminUser)
		form.Password = utils.GetMd5String(form.Username+form.Password)
		_ = rpccli.Call("Login", form, res)
		if res.Adminid > 0 {
			//session := sessions.Default(context)
			//session.Set(utils.ASYUSERID, res)
			//token := utils.Random62String(64)
			//context.Header(utils.ASYUSERID, res.Adminname)
			//context.Header(utils.ASTROLE, res.Rolename)
			//context.Header(utils.ASYTOKEN, token)
			//context.SetCookie(utils.ASYUSERID, res.Adminname, 3600, "/", config.Get("domain").String("localhost"), false, true)
			//context.SetCookie(utils.ASYUSERID, res.Rolename, 3600, "/", config.Get("domain").String("localhost"), false, true)
			//context.SetCookie(utils.ASYTOKEN, token, 3600, "/", config.Get("domain").String("localhost"), false, true)
			//	//err = redis.Set(result.UserId, token)
			//	//utils.WriteErrorLog(ctx.FullPath(), err)
			//context.HTML(http.StatusOK, "admin/index/index.tmpl", gin.H{
			//	"status": "you are logged in",
			//	"res":res,
			//	"form":form,
			//})
			sessionKey := "Admin"
			session := sessions.Default(context)
			session.Set(sessionKey, res)
			session.Save()
			context.Redirect(http.StatusMovedPermanently, "/index")
			return
		}
		context.HTML(http.StatusOK, "admin/index/login.tmpl", gin.H{
			"MXCMS_SOFTNAME":"MXCMS",
		})
		return
	}
	context.HTML(http.StatusOK, "admin/index/login.tmpl", gin.H{
		"MXCMS_SOFTNAME":"MXCMS",
	})
	return
}