package service

import (
	"mxcms/app"
	"mxcms/app/models"
)

func Login(username string, password string) (*models.AdminUser) {
	var result models.AdminUser
	app.Db.Table("mx_admin").Where("adminname=? and password=?",username,password).First(&result)
	return &result
}

func GetMenu(where string, args ...interface{}) (*[]models.Menu) {
	var result []models.Menu
	app.Db.Where(where, args...).Find(&result)
	return &result
}

func GetAdminRolePriv(where string, args ...interface{}) (*[]models.AdminRolePriv) {
	var result []models.AdminRolePriv
	app.Db.Where(where, args...).Find(&result)
	return &result
}