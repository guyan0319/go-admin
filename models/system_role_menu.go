package models

import "fmt"

type SystemRoleMenu struct {
	Id           int `json:"id" xorm:"not null pk autoincr comment('主键') INT(11)"`
	SystemRoleId int `json:"system_role_id" xorm:"not null default 0 comment('角色主键') index(system_role_id) INT(11)"`
	SystemMenuId int `json:"system_menu_id" xorm:"not null default 0 comment('菜单主键') index(system_role_id) INT(11)"`
}
var systemrolemenu = "system_role_menu"
func (rm *SystemRoleMenu) Add() bool{
	if rm.SystemRoleId==0 || rm.SystemMenuId==0 {
		return false
	}
	rolemenu :=SystemRoleMenu{}
	has, err := mEngine.Where("system_role_id = ?", rm.SystemRoleId).Where("system_menu_id=?",rm.SystemMenuId).Get(&rolemenu)
	if err!=nil {
		fmt.Println(err)
		return false
	}
	if has {
		return true
	}
	_,err =mEngine.Insert(rm)
	if err!=nil {
		fmt.Println(err)
		return false
	}
	return true
}

func (rm *SystemRoleMenu) GetAll()([]SystemRoleMenu,error) {
	var systemrolemenus []SystemRoleMenu
	err:=mEngine.Find(&systemrolemenus)
	return systemrolemenus,err
}