package systemnewdb

var menuStateOk = 1
var menuStateNo = 0

func (m *SystemMenu) GetAll() ([]SystemMenu, error) {
	var systemmenus []SystemMenu
	result := db.Find(&systemmenus)
	return systemmenus, result.Error
}
func (m *SystemMenu) GetRowByUid(uid interface{}) []SystemMenu {
	var menu []SystemMenu
	db.Model(m).Distinct("system_menu.*").Joins("INNER JOIN system_role_menu ON system_role_menu.system_menu_id=system_menu.id").
		Joins("INNER JOIN system_user_role ON system_user_role.system_role_id=system_role_menu.system_role_id").
		Where("system_menu.state=?", m.State).
		Where("system_user_role.system_user_id=?", uid).
		Find(&menu)
	return menu
}
func (m *SystemMenu) GetRouteByRole(id interface{}) []SystemMenu {
	var constant []SystemMenu
	menu := SystemMenu{Type: 1}
	constant, _ = menu.GetRowByType()

	var end []SystemMenu
	menu.Type = 3
	end, _ = menu.GetRowByType()
	var async []SystemMenu
	async = menu.GetRowByRole(id)
	constant = append(constant, async...)
	constant = append(constant, end...)
	return constant
}
func (m *SystemMenu) GetRowByType() ([]SystemMenu, error) {
	var systemmenus []SystemMenu
	result := db.Where("type=?", m.Type).Find(&systemmenus)
	return systemmenus, result.Error
}
func (m *SystemMenu) GetRowByRole(id interface{}) []SystemMenu {
	var menu []SystemMenu
	db.Model(m).Select("system_menu.*").
		Joins("INNER JOIN system_role_menu ON system_role_menu.system_menu_id=system_menu.id").
		Where("system_menu.state=?", 1).
		Where("system_role_menu.system_role_id=?", id).
		Find(&menu)
	return menu
}

//根据主键获取一条数据

func (m *SystemMenu) GetRow() int64 {
	result := db.First(&m)
	return result.RowsAffected
}

func (m *SystemMenu) Update(menu SystemMenu) int64 {
	result := db.Model(&m).Updates(menu)
	return result.RowsAffected
}

func (m *SystemMenu) Delete() error {
	result := db.Model(&m).Where("id=?", m.ID).Update("state", menuStateNo)
	return result.Error
}
func (m *SystemMenu) GetRowByPathCT(menu SystemMenu) bool {
	res := db.Where("path=?", menu.Path).Where("component=?", menu.Component).Where("type=?", menu.Type).First(&m)
	if res.Error == nil && res.RowsAffected > 0 {
		return true
	}
	return false
}
