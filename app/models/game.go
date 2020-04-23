package models

type GameReview struct{
	GameId string `json:"gameid" gorm:"primary_key"`
	Name string  `json:"name" gorm:"type:varchar(20);not null;"`
	Avatar string `json:"avatar" gorm:"type:text;not null;"`
	Platform string `json:"platform" gorm:"type:varchar(128);not null;"`
	Good string `json:"good" gorm:"type:text;not null;"`
	Bad string `json:"Bad" gorm:"type:text;not null;"`
	Score float32 `json:"score" gorm:"type:int(8);not null;"`
	Scoreword string `json:"scoreword" gorm:"type:varchar(20);not null;"`
}

type Menu struct {
	Id int `gorm:"smallint(6) unsigned NOT NULL AUTO_INCREMENT,"`
	Name string `gorm:"char(40) NOT NULL DEFAULT '',"`
	Parentid int `gorm:"smallint(6) NOT NULL DEFAULT '0',"`
	M string `gorm:"char(20) NOT NULL DEFAULT '',"`
	C string `gorm:"char(20) NOT NULL DEFAULT '',"`
	A string `gorm:"char(30) NOT NULL DEFAULT '',"`
	Data string `gorm:"char(100) NOT NULL DEFAULT '',"`
	Listorder int `gorm:"smallint(6) unsigned NOT NULL DEFAULT '0',"`
	Display int `gorm:"tinyint(1) NOT NULL DEFAULT '0',"`
}

type AdminRolePriv struct {
	Roleid int `gorm:tinyint(3) unsigned NOT NULL DEFAULT '0',`
	M string `gorm:char(20) NOT NULL DEFAULT '',`
	C string `gorm:char(20) NOT NULL DEFAULT '',`
	A string `gorm:char(30) NOT NULL DEFAULT '',`
	Data string `gorm:char(100) NOT NULL DEFAULT '',`
}

