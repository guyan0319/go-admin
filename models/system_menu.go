package models

import (
	"time"
)

type SystemMenu struct {
	Id     int       `json:"id" xorm:"not null pk autoincr comment('主键') INT(11)"`
	Name   string    `json:"name" xorm:"not null comment('名称') VARCHAR(100)"`
	Type   string    `json:"type" xorm:"not null comment('类型(menu,button)') VARCHAR(100)"`
	Url    string    `json:"url" xorm:"not null comment('url') VARCHAR(200)"`
	Nav    int       `json:"nav" xorm:"not null default 1 comment('是否在导航显示（0不显示1显示）') TINYINT(4)"`
	Icon   string    `json:"icon" xorm:"not null default '' comment('菜单图标') VARCHAR(100)"`
	Target string    `json:"target" xorm:"not null default '_self' comment('打开方式') VARCHAR(20)"`
	Params string    `json:"params" xorm:"not null default '' comment('链接参数') VARCHAR(500)"`
	Pid    int       `json:"pid" xorm:"not null comment('父ID') index(idx_list) INT(11)"`
	Sort   int       `json:"sort" xorm:"not null default 0 comment('排序') index(idx_list) INT(11)"`
	Status int       `json:"status" xorm:"not null default 1 comment('状态（0禁止1启动）') index(idx_list) TINYINT(4)"`
	Level  int       `json:"level" xorm:"not null comment('层级') TINYINT(4)"`
	Ctime  time.Time `json:"ctime" xorm:"not null comment('时间') DATETIME"`
}
