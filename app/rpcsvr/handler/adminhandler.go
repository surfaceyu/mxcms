package handler

import (
"context"
"mxcms/app/models"
"mxcms/app/rpcsvr/service"
)

func (this *Handlers) Login(ctx context.Context, req *models.Login, rsp *models.AdminUser) error {
	result, err := service.Login(req.Username, req.Password)
	*rsp = *result
	return err
}
