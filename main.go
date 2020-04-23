package main

import (
	//"mxcms/app/admin"
	//"mxcms/app/web"
	//"mxcms/app/rpcsvr"
	"mxcms/app/admin"
	"mxcms/app/rpcsvr"
)
func main() {
	go rpcsvr.Start()
	//go web.Start()
	admin.Start()
	//go srv.Start()
}