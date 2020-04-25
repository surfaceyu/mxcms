package app

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/source/file"
	"log"
	"mxcms/mxcore/databases/mysql"
	"mxcms/mxcore/databases/redis"
)

const (
	defaultConfigPath = "app/config.json"
)

var Db *gorm.DB

func loadConfig() {
	if err := config.Load(file.NewSource(
		file.WithPath(defaultConfigPath),
	)); err != nil {
		log.Panic(err)
	}
	fmt.Println("loadConfig", config.Map())
}

func init() {
	Appinit()
}

func Appinit() {
	//fmt.Println("init databases")
	loadConfig()
	Db = mysql.InitMysql()
	Db.SingularTable(true)

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "mx_" + defaultTableName
	}

	redis.InitRedis()
}
