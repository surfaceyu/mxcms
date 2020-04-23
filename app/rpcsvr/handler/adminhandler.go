package handler

import (
	"context"
	"mxcms/app/models"
	"mxcms/app/rpcsvr/service"
)

func (this *Handlers) Login(ctx context.Context, req *models.Login, rsp *models.AdminUser) error {
	result := service.Login(req.Username, req.Password)
	*rsp = *result
	return nil
}

func (this *Handlers) GetMenu(ctx context.Context, req *models.Menu, rsp *[]models.Menu) error {
	result:= service.GetMenu("")
	*rsp = *result
	return nil
}

func (this *Handlers) GetAdminRolePriv(ctx context.Context, roleId int, rsp *[]models.AdminRolePriv) error {
	result := service.GetAdminRolePriv("roleid=?", roleId)
	*rsp = *result
	return nil
}

func (this *Handlers) GetAdminMenu(ctx context.Context, roleId int, rsp *[]models.Msgmenu) error {
	parentMenus:= service.GetMenu("parentid=? and display=?", 0,1)
	for _, value := range *parentMenus{
		childs:= service.GetMenu("parentid=? and display=?", value.Id,1)
		rolepriv := service.GetAdminRolePriv("roleid=?", roleId)
		for k:=0; k<len(*childs); k++{
			v := (*childs)[k]
			if roleId != 1 {
				find := false
				for  _, priv := range *rolepriv{
					if priv.A == v.A && priv.C == v.C && priv.M == v.M {
						find = true
					}
				}
				if !find {
					//删除
					*childs = append((*childs)[:k], (*childs)[k+1:]...)
					k--
				}
			}
		}
		if len(*childs) >= 1 {
			*rsp = append(*rsp, models.Msgmenu{
				Menu:value,
				Child:*childs,
			})
		}
	}
	return nil
}