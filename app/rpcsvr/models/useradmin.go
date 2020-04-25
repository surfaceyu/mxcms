package models

import (
	"github.com/jinzhu/gorm"
	"mxcms/app"
	"mxcms/app/models"
)

func Login(username string, password string) (*models.AdminUser, error) {
	var result models.AdminUser
	err := app.Db.Table("mx_admin").Where("adminname=? and password=?",username,password).First(&result).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &result, nil
}

// ExistAdminByName checks if an admin exists based on ID
func ExistAdminByName(adminname string) (bool, error) {
	var adminUser models.AdminUser
	err := app.Db.Table("mx_admin").Select("adminid").Where("adminname = ?", adminname).First(&adminUser).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if adminUser.Adminid > 0 {
		return true, nil
	}

	return false, nil
}

// GetUserAdmin Get UserAdmin based on username password
func GetUserAdmin(username string, password string) (*models.AdminUser, error) {
	var result models.AdminUser
	err := app.Db.Table("mx_admin").Where("adminname=? and password=?",username,password).First(&result).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &result, nil
}
