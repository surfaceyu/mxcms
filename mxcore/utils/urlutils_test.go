package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"testing"
)

func TestStart(t *testing.T) {
	var RouteInfos []gin.RouteInfo
	RouteInfos = append(RouteInfos, gin.RouteInfo{
		Method:"GET",
		Path:"/login",
		Handler:"mxcms/app/admin/controller.Login",
	})
	RouteInfos = append(RouteInfos, gin.RouteInfo{
		Method:"POST",
		Path:"/login",
		Handler:"mxcms/app/admin/controller.Login",
	})
	RouteInfos = append(RouteInfos, gin.RouteInfo{
		Method:"GET",
		Path:"/index",
		Handler:"mxcms/app/admin/controller.Index",
	})
	RouteInfos = append(RouteInfos, gin.RouteInfo{
		Method:"GET",
		Path:"/index/",
		Handler:"mxcms/app/admin/controller.Index",
	})
	s1:=URLFor(RouteInfos)("Logout", "name", "astaxie", "age", "25")
	s2:=URLFor(RouteInfos)("Login", ":name", "astaxie", ":age", "25")
	s3:=URLFor(RouteInfos)("Login")
	fmt.Println(s1, s2, s3)
}

