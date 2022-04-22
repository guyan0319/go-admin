package systemnewdb

func (ur *SystemUserRole) GetRowByUid() []string {
	var role []string
	db.Model(&ur).Select("system_role.name").Joins("left join system_role on system_role.id=system_user_role.system_role_id").Where("system_user_id=?", ur.SystemUserID).Scan(&role)
	return role
}
