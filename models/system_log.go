package models

import (
	"time"
)

type SystemLog struct {
	Id            int64     `json:"id" xorm:"pk autoincr comment('主键') BIGINT(20)"`
	SystemUserId  int       `json:"system_user_id" xorm:"default 0 comment('主键') index INT(11)"`
	Title         string    `json:"title" xorm:"not null default '' comment('日志标题') VARCHAR(300)"`
	Content       string    `json:"content" xorm:"comment('日志内容记录SQL') TEXT"`
	RelationId    int64     `json:"relation_id" xorm:"not null default 0 comment('相关对应表主键') index BIGINT(20)"`
	RelationTable int       `json:"relation_table" xorm:"not null default 1 comment('对应表(1 system_user,2 system_menu,3 system_role)') index INT(4)"`
	Ip            string    `json:"ip" xorm:"not null default '' comment('ip') VARCHAR(50)"`
	Url           string    `json:"url" xorm:"not null default '' VARCHAR(500)"`
	Ctime         time.Time `json:"ctime" xorm:"not null default '0000-00-00 00:00:00' comment('时间') index DATETIME"`
}
