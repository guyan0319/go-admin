package models

import (
	"fmt"
	"time"
)

type SystemRole struct {
	Id          int       `json:"id" xorm:"not null pk autoincr comment('主键') INT(11)"`
	Name        string    `json:"name" xorm:"not null comment('角色名称') VARCHAR(100)"`
	AliasName   string    `json:"alias_name" xorm:"not null default '' comment('别名') VARCHAR(50)"`
	Description string    `json:"description" xorm:"not null default '' comment('描述') VARCHAR(200)"`
	Status      int       `json:"status" xorm:"not null default 1 comment('角色状态（0无效1有效）') index TINYINT(4)"`
	Type        int       `json:"type" xorm:"not null default 1 comment('属于哪个应用') index INT(4)"`
	Ctime       time.Time `json:"ctime" xorm:"not null comment('创建时间') DATETIME"`
}

var systemrole = "system_role"

func (r *SystemRole) GetRowById() bool {
	has, err := mEngine.Where("id = ?", r.Id).Get(&r)
	if err == nil && has {
		return true
	}
	return false
}
func (r *SystemRole) GetRowByName() bool {
	role := SystemRole{}
	has, err := mEngine.Where("name = ?", r.Name).Get(&role)
	if err == nil && has {
		return true
	}
	return false
}
func (r *SystemRole) Add() bool {
	if r.Name == "" {
		return false
	}
	role := SystemRole{}
	has, err := mEngine.Where("name = ?", r.Name).Desc("id").Get(&role)
	if err != nil {
		fmt.Println(err)
		return false
	}
	if has {
		r.Id=role.Id
		return true
	}
	r.Status=1
	r.Type=1
	_, err = mEngine.Insert(r)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
