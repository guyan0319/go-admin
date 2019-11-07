package models

type SystemUserRole struct {
	Id           int `json:"id" xorm:"not null pk autoincr comment('主键') INT(11)"`
	SystemUserId int `json:"system_user_id" xorm:"not null comment('用户主键') index(system_user_id) INT(11)"`
	SystemRoleId int `json:"system_role_id" xorm:"not null comment('角色主键') index(system_user_id) INT(11)"`
}
var systemuserrole ="system_user_role"

func(u *SystemUserRole) GetRow() bool {
	has, err := mEngine.Get(u)
	if err==nil &&  has  {
		return true
	}
	return false
}
func(u *SystemUserRole) GetRowByUid() ([]string,error) {
	var role []string
	err := mEngine.Table(systemuserrole).Select(systemrole+".name").
		Join("INNER", systemrole, systemuserrole+".system_role_id = "+systemrole+".id").
		Where(systemrole+".status = ?", 1).
		Where(systemuserrole+".system_user_id = ?", u.SystemUserId).
		Find(&role)
	return role,err
}
