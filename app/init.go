package app

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/source/file"
	"log"
	"mxcms/mxcore/databases/mysql"
)

const (
	defaultConfigPath = "app/config.json"
)

var Db *gorm.DB

func init() {
	fmt.Println("init databases")
	loadConfig()
	Db = mysql.InitMysql()
	//databases.InitRedis()
}

func loadConfig() {
	if err := config.Load(file.NewSource(
		file.WithPath(defaultConfigPath),
	)); err != nil {
		log.Panic(err)
	}
	fmt.Println("loadConfig", config.Map())
}
