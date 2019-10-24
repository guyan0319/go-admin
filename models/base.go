package models

import (
	"github.com/xormplus/xorm"
	"go-admin/conf"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strconv"
)


var mEngine *xorm.Engine

func init() {
	if mEngine == nil {
		var err error
		mEngine, err = xorm.NewEngine(conf.Db["db1"]["driverName"], conf.Db["db1"]["dsn"])
		if err != nil {
			log.Fatal(err)
		}
		n, _ := strconv.Atoi(conf.Db["db1"]["maxIdle"])

		mEngine.SetMaxIdleConns(n) //空闲连接
		n, _ = strconv.Atoi(conf.Db["db1"]["maxOpen"])
		mEngine.SetMaxOpenConns(n) //最大连接数
		mEngine.ShowSQL(true)
		mEngine.ShowExecTime(true)
	}
}

