package models

type SystemUserRole struct {
	Id           int `json:"id" xorm:"not null pk autoincr comment('主键') INT(11)"`
	SystemUserId int `json:"system_user_id" xorm:"not null comment('用户主键') index(system_user_id) INT(11)"`
	SystemRoleId int `json:"system_role_id" xorm:"not null comment('角色主键') index(system_user_id) INT(11)"`
}

