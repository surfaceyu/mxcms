package mysql

import (
	//"database/sql"
	//"log"
	//"os"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/micro/go-micro/v2/config"
)

func InitMysql() *gorm.DB {
	//db, err := sql.Open("mysql", config.Get("mysql", "hrefs").String(""))
	//checkErr(err, "sql.Open failed")
	//
	//db.SetMaxIdleConns(config.Get("mysql", "MaxIdleConns").Int(5))
	//db.SetMaxOpenConns(config.Get("mysql", "MaxOpenConns").Int(50))
	//dbMap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{}}
	//if config.Get("mysql", "gorp", "trace-on").Bool(false) {
	//	dbMap.TraceOn("[gorp]", log.New(os.Stdout, "[SQL]:", log.Lmicroseconds))
	//}
	mysqlClient, err := gorm.Open("mysql", config.Get("databases", "mysql", "hrefs").String(""))
	if err != nil {
		panic("连接数据库失败")
	}

	return mysqlClient
}
