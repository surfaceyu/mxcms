package models

type Login struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type AdminUser struct {
	Adminid int `gorm:"MEDIUMINT(6) UNSIGNED NOT NULL AUTO_INCREMENT"`
	Adminname string `gorm:"VARCHAR(30) NOT NULL DEFAULT ''"`
	Password string `gorm:"VARCHAR(32) NOT NULL DEFAULT ''"`
	Roleid int `gorm:"TINYINT(3) UNSIGNED NOT NULL DEFAULT '1'"`
	Rolename string `gorm:"VARCHAR(30) NOT NULL DEFAULT ''"`
	Realname string `gorm:"VARCHAR(30) NOT NULL DEFAULT ''"`
	Nickname string `gorm:"VARCHAR(30) NOT NULL DEFAULT ''"`
	Email string `gorm:"VARCHAR(30) NOT NULL DEFAULT ''"`
	Logintime int `gorm:"INT(10) UNSIGNED NOT NULL DEFAULT '0'"`
	Loginip string `gorm:"VARCHAR(15) NOT NULL DEFAULT ''"`
	Addtime string `gorm:"INT(10) UNSIGNED NOT NULL DEFAULT '0'"`
	Addpeople string `gorm:"VARCHAR(30) NOT NULL DEFAULT ''"`
}