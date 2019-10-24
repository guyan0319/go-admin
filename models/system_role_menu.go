package models

type SystemRoleMenu struct {
	Id           int `json:"id" xorm:"not null pk autoincr comment('主键') INT(11)"`
	SystemRoleId int `json:"system_role_id" xorm:"not null default 0 comment('角色主键') index(system_role_id) INT(11)"`
	SystemMenuId int `json:"system_menu_id" xorm:"not null default 0 comment('菜单主键') index(system_role_id) INT(11)"`
}
