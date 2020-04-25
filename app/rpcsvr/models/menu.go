package models

import (
	"github.com/jinzhu/gorm"
	"mxcms/app"
	"mxcms/app/models"
)

func GetMenu(where string, args ...interface{}) (*[]models.Menu, error) {
	var result []models.Menu
	err := app.Db.Where(where, args...).Order("listorder ASC").Find(&result).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &result, nil
}

func GetAdminRolePriv(where string, args ...interface{}) (*[]models.AdminRolePriv, error) {
	var result []models.AdminRolePriv
	err := app.Db.Where(where, args...).Find(&result).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &result, nil
}
