package systemnewdb

import (
	"go-admin/lib/common"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
	"log"
	"os"
	"time"
)

var db *gorm.DB

func init() {
	c := common.Conf
	var err error
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,        // 禁用彩色打印
		},
	)
	db, err = gorm.Open(mysql.Open(c.Db["systemnewdb"].Dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true}, //禁止表名后加s，即全局禁用表名复数
		Logger:         newLogger,
		//Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalf("error: database connect  %s", err.Error())
	}
	db.Use(dbresolver.Register(dbresolver.Config{
		Replicas: []gorm.Dialector{mysql.Open(c.Db["systemnewdb"].Dsn)},
	}, "systemnewdb").
		SetConnMaxIdleTime(time.Hour).
		SetConnMaxLifetime(24 * time.Hour).
		SetMaxIdleConns(100).
		SetMaxOpenConns(200))

}
func GetDb() *gorm.DB {
	return db
}
