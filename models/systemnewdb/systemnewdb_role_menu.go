package systemnewdb

func (rm *SystemRoleMenu) GetAll() ([]SystemRoleMenu, error) {
	var systemrolemenus []SystemRoleMenu
	result := db.Find(&systemrolemenus)
	return systemrolemenus, result.Error
}
