package rpcsvr

import (
	"context"
	"github.com/davyxu/golog"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/server"
	"mxcms/app/rpcsvr/handler"
	"mxcms/mxcore/discovery"
	"time"
)

var glog *golog.Logger = golog.New("mx-rpc-server")

func Start() {
	urls := discovery.GetEtcdUrls()
	reg := etcd.NewRegistry(func(op *registry.Options) {
		op.Addrs = urls
	})

	service := micro.NewService(
		micro.Registry(reg),
		micro.Name(config.Get("srv").String("micro.qgn.srv")),
		micro.WrapHandler(logWrapper),
	)
	server.Init()
	service.Server().Init(server.Wait(nil))
	micro.RegisterHandler(service.Server(), new(handler.Handlers))
	service.Run()
}

func logWrapper(fn server.HandlerFunc) server.HandlerFunc {
	start := time.Now()
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		err := fn(ctx, req, rsp)
		if err != nil {
			glog.Errorf("%s %s", time.Since(start), req.Endpoint())
		}
		glog.Infof("%s %s", time.Since(start), req.Endpoint())
		return err
	}
}
