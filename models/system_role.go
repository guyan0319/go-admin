package models

import (
	"time"
)

type SystemRole struct {
	Id     int       `json:"id" xorm:"not null pk autoincr comment('主键') INT(11)"`
	Name   string    `json:"name" xorm:"not null comment('角色名称') VARCHAR(100)"`
	Status int       `json:"status" xorm:"not null default 1 comment('角色状态（0无效1有效）') index TINYINT(4)"`
	Type   int       `json:"type" xorm:"not null default 1 comment('属于哪个应用') index INT(4)"`
	Ctime  time.Time `json:"ctime" xorm:"not null comment('创建时间') DATETIME"`
}

func(r *SystemRole) GetRowById() bool {
	has, err := mEngine.Where("id = ?", r.Id).Get(r)
	if err==nil &&  has  {
		return true
	}
	return false
}

