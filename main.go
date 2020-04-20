package main

import (
	"mxcms/app/admin"
	"mxcms/app/web"
	"mxcms/app/rpcsvr"
)
func main() {
	go rpcsvr.Start()
	go admin.Start()
	web.Start()
	//go srv.Start()
}