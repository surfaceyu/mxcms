package service

import (
	"encoding/json"
	"errors"
	"mxcms/app/models"
	mxcache "mxcms/app/rpcsvr/cache"
	mxconst "mxcms/app/rpcsvr/common"
	m "mxcms/app/rpcsvr/models"
	mxredis "mxcms/mxcore/databases/redis"
)

type Menu models.Menu
type AdminRolePriv models.AdminRolePriv

func GetAdminMenu(roleId int, rsp *[]models.Msgmenu) (error){
	var parentMenuService Menu
	var RoleprivService AdminRolePriv

	var cache mxcache.Menu
	key := cache.GetAdminKey(roleId)
	if mxredis.Exists(key) {
		data, err := mxredis.Get(key)
		if err != nil {
			mxlog.Infoln(err)
		} else {
			json.Unmarshal(data, &rsp)
			return nil
		}
	}

	parentMenus, err := parentMenuService.Get(0,"parentid=? and display=?", 0,1)
	if err != nil {
		mxlog.Infoln(mxconst.ERROR_GET_MENU_FAIL)
		return errors.New(mxconst.ERROR_GET_MENU_FAIL)
	}
	for _, value := range *parentMenus{
		childs, err:= parentMenuService.Get(value.Id,"parentid=? and display=?", value.Id,1)
		if err != nil {
			mxlog.Infoln(mxconst.ERROR_GET_MENU_FAIL)
			return errors.New(mxconst.ERROR_GET_MENU_FAIL)
		}

		rolepriv, err := RoleprivService.Get(roleId)

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
	mxredis.Set(key, *rsp, 3600)
	return nil
}

func (this *Menu)Get(parentId int, where string, args ...interface{}) (*[]models.Menu, error) {
	var cacheMenu *[]models.Menu
	var cache mxcache.Menu

	key := cache.GetKey(parentId)
	if mxredis.Exists(key) {
		data, err := mxredis.Get(key)
		if err != nil {
			mxlog.Infoln(err)
		} else {
			json.Unmarshal(data, &cacheMenu)
			return cacheMenu, nil
		}
	}

	result, err := m.GetMenu(where, args...)
	if err != nil {
		return nil, err
	}

	mxredis.Set(key, result, 3600)
	return result, nil
}

func (this *AdminRolePriv) Get(roleId int) (*[]models.AdminRolePriv, error) {
	var cacheMenu *[]models.AdminRolePriv
	var cache mxcache.AdminRolePriv

	key := cache.GetKey(roleId)
	if mxredis.Exists(key) {
		data, err := mxredis.Get(key)
		if err != nil {
			mxlog.Infoln(err)
		} else {
			json.Unmarshal(data, &cacheMenu)
			return cacheMenu, nil
		}
	}

	result, err := m.GetAdminRolePriv("roleid=?", roleId)
	if err != nil {
		return nil, err
	}

	mxredis.Set(key, result, 3600)
	return result, nil
}
