package handler

import (
	"errors"
	"context"
	"github.com/davyxu/golog"
	"mxcms/app/models"
	mxconst "mxcms/app/rpcsvr/common"
	"mxcms/app/rpcsvr/service"
)

var mxlog *golog.Logger = golog.New("mx-service-adminhandld")

func (this *Handlers) Login(ctx context.Context, req *models.MsgLogin, rsp *models.AdminUser) error {
	adminUserService := service.AdminUser{
		Adminname: req.Username,
		Password: req.Password,
	}
	exists, err := adminUserService.ExistByName()
	if err != nil {
		mxlog.Infoln(mxconst.ERROR_GET_USER_ADMIN_FAIL)
		return errors.New(mxconst.ERROR_GET_USER_ADMIN_FAIL)
	}
	if !exists {
		mxlog.Infoln(mxconst.ERROR_NOT_EXIST_USER_ADMIN)
		return errors.New(mxconst.ERROR_NOT_EXIST_USER_ADMIN)
	}

	result, err := adminUserService.Get()
	if err != nil {
		mxlog.Infoln(mxconst.ERROR_GET_USER_ADMIN_FAIL)
		return errors.New(mxconst.ERROR_GET_USER_ADMIN_FAIL)
	}
	//result := service.Login(req.Username, req.Password)
	*rsp = *result
	return nil
}

func (this *Handlers) GetAdminMenu(ctx context.Context, roleId int, rsp *[]models.Msgmenu) error {
	service.GetAdminMenu(roleId, rsp)
	return nil
}