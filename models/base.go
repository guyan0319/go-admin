package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"go-admin/conf"
	"log"
)
var mEngine *xorm.Engine

func init() {
	if mEngine == nil {
		var err error
		mEngine, err = xorm.NewEngine(conf.Db["db1"].DriverName ,conf.Db["db1"].Dsn)
		if err != nil {
			log.Fatal(err)
		}
		mEngine.SetMaxIdleConns(conf.Db["db1"].MaxIdle) //空闲连接
		mEngine.SetMaxOpenConns(conf.Db["db1"].MaxOpen) //最大连接数
		mEngine.ShowSQL(conf.Db["db1"].ShowSql)
		mEngine.ShowExecTime(conf.Db["db1"].ShowExecTime)
	}
}

