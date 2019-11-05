package models

import (
	"time"
)

type SystemMenu struct {
	Id          int       `json:"id" xorm:"not null pk autoincr comment('主键') INT(11)"`
	Name        string    `json:"name" xorm:"not null default '' comment('名称') VARCHAR(100)"`
	Path        string    `json:"path" xorm:"not null default '' comment('路径') index VARCHAR(50)"`
	Component   string    `json:"component" xorm:"not null default '' comment('组件') VARCHAR(100)"`
	Redirect    string    `json:"redirect" xorm:"not null default '' comment('重定向') VARCHAR(200)"`
	MetaTitle   string    `json:"meta_title" xorm:"not null default '' comment('meta标题') VARCHAR(50)"`
	MetaIcon    string    `json:"meta_icon" xorm:"not null default '' comment('meta icon') VARCHAR(50)"`
	MetaNocache int       `json:"meta_nocache" xorm:"not null default 0 comment('是否缓存（1:是 0:否）') TINYINT(4)"`
	Alwaysshow  int       `json:"alwaysshow" xorm:"not null default 0 comment('是否总是显示（1:是0：否）') TINYINT(4)"`
	Type        int       `json:"type" xorm:"not null default 2 comment('类型(1:固定,2:权限配置)') TINYINT(4)"`
	Hidden      int       `json:"hidden" xorm:"not null default 1 comment('是否隐藏（0否1是）') TINYINT(4)"`
	Pid         int       `json:"pid" xorm:"not null default 0 comment('父ID') index(idx_list) INT(11)"`
	Sort        int       `json:"sort" xorm:"not null default 0 comment('排序') index(idx_list) INT(11)"`
	Status      int       `json:"status" xorm:"not null default 1 comment('状态（0禁止1启动）') index(idx_list) TINYINT(4)"`
	Level       int       `json:"level" xorm:"not null default 0 comment('层级') TINYINT(4)"`
	Ctime       time.Time `json:"ctime" xorm:"not null default '0000-00-00 00:00:00' comment('时间') DATETIME"`
}
