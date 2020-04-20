package middleware

import (
	"github.com/davyxu/golog"
	"github.com/gin-gonic/gin"
	"time"
)

var glog *golog.Logger = golog.New("mx-trace")

func TraceGin(ctx *gin.Context) {
	defer trace(ctx.Request.Method, ctx.Request.URL.Path)()
	ctx.Next()
}

func trace(method string, path string) func() {
	start := time.Now()
	return func() {
		glog.Infof("%s %s %s", time.Since(start), method, path)
	}
}
