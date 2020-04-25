package service

import (
	"encoding/json"
	"mxcms/app"
	"mxcms/app/models"
	mxcache "mxcms/app/rpcsvr/cache"
	m "mxcms/app/rpcsvr/models"
	mxredis "mxcms/mxcore/databases/redis"

	golog "github.com/davyxu/golog"
)

type AdminUser models.AdminUser

var mxlog *golog.Logger = golog.New("mx-service-useradmin")

func (this *AdminUser) Login(username string, password string) (*models.AdminUser) {
	var result models.AdminUser
	app.Db.Table("mx_admin").Where("adminname=? and password=?",username,password).First(&result)
	return &result
}

func (this *AdminUser) ExistByName() (bool, error) {
	return m.ExistAdminByName(this.Adminname)
}

func (this *AdminUser) Get() (*models.AdminUser, error) {
	var cacheAdminUser *models.AdminUser

	cache := &mxcache.AdminUser{
		Adminname: this.Adminname,
		Password: this.Password,
	}
	key := cache.GetKey()
	if mxredis.Exists(key) {
		data, err := mxredis.Get(key)
		if err != nil {
			mxlog.Infoln(err)
		} else {
			json.Unmarshal(data, &cacheAdminUser)
			return cacheAdminUser, nil
		}
	}

	result, err := m.GetUserAdmin(this.Adminname, this.Password)
	if err != nil {
		return nil, err
	}

	mxredis.Set(key, result, 3600)
	return result, nil
}

//func Login(username string, password string) (*models.AdminUser) {
//	var result models.AdminUser
//	app.Db.Table("mx_admin").Where("adminname=? and password=?",username,password).First(&result)
//	return &result
//}