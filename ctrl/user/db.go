package user

import (
	"fmt"
	"gorm.io/plugin/dbresolver"
	log2 "log"
	"web-demo/models/systemdb"
)

//
//gorm 数据库操作
//https://gorm.io/zh_CN/docs/sql_builder.html
func Login() {
	log2.Fatalf("fasdfadsf")
	db := systemdb.GetDb()
	logMgr := systemdb.SystemLogMgr(db)
	row, err := logMgr.FetchByPrimaryKey(2)
	if err != nil {

	}
	tx, err := logMgr.Gets()
	for _, v := range tx {
		fmt.Println(v.ID)
	}
	db1 := db.Clauses(dbresolver.Use("systemrdb")) //切换只读
	logMgrs := systemdb.SystemLogMgr(db1)
	row, err = logMgrs.FetchByPrimaryKey(2)
	if err != nil {

	}
	fmt.Println(row)
	log := systemdb.SystemLog{}
	db.Find(&log) //使用只读systemrmdb库
	fmt.Println(log)

	//var id int64
	//for ret.Next() {
	//	ret.Scan(&id)
	//	fmt.Println(id)
	//	// do something
	//}
	//fmt.Println( ret,"fasdf")
	// Global Resolver 示例
	//tx,err:=db.Find(&systemdb.SystemLog{}).Rows() // replicas `db3`/`db4`

	//插入
	//log:=systemdb.SystemLog{
	//	Ctime: time.Now(),
	//}
	//
	//result:=db.Create(&log)
	//fmt.Println(result)
	//fmt.Println(row)

}
