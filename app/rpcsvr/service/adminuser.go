package service

import (
	"mxcms/app"
	"mxcms/app/models"
)

func Login(username string, password string) (*models.AdminUser, error) {
	var result models.AdminUser
	app.Db.Table("mx_admin").Where("adminname=? and password=?",username,password).First(&result)
	return &result, nil
}