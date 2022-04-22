package systemnewdb

import (
	"go-admin/lib/common"
)

type SearchUser struct {
	Name string `json:"name" xorm:"not null comment('姓名') VARCHAR(50)"`
}

//根据非零值获取一条数据
func (u *SystemUser) GetRow() (int64, error) {
	result := db.First(&u)
	return result.RowsAffected, result.Error
}
func (u *SystemUser) GetAllPage(paging *common.Paging) ([]SystemUser, error) {
	var systemusers []SystemUser
	var err error
	var count int64
	db.Model(&u).Where("state=?", 1).Count(&count)
	paging.Total = count
	paging.GetPages()
	if paging.Total < 1 {
		return systemusers, err
	}
	result := db.Where("state=?", 1).Limit(int(paging.PageSize)).Offset(int(paging.StartNums)).Find(&systemusers)
	return systemusers, result.Error
}

func (u *SystemUser) GetAllByName() ([]SearchUser, error) {
	var systemusers []SearchUser
	res := db.Model(&u).Where("name like ?", u.Name+"%").Select("name").Find(&systemusers)
	return systemusers, res.Error
}

func (u *SystemUser) Add(roles []interface{}) error {
	// 再唠叨一下，事务一旦开始，你就应该使用 tx 处理数据
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}
	//Create 会更新u主键
	if err := tx.Create(u).Error; err != nil {
		tx.Rollback()
		return err
	}
	//如果没有设置权限
	if len(roles) < 1 {
		return tx.Commit().Error
	}
	for _, k := range roles {
		roleModel := SystemRole{Name: k.(string)}
		res := tx.First(&roleModel)
		if res.RowsAffected < 1 || roleModel.State == 0 {
			continue
		}
		userroleModel := SystemUserRole{SystemRoleID: roleModel.ID, SystemUserID: u.ID}
		res = tx.First(&userroleModel)
		if res.RowsAffected > 0 {
			continue
		}
		if err := tx.Create(&userroleModel).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit().Error
}

func (u *SystemUser) Update(roles []interface{}) error {
	// 再唠叨一下，事务一旦开始，你就应该使用 tx 处理数据
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Model(&SystemUser{}).Where("id = ?", u.ID).Updates(&u).Error; err != nil {
		tx.Rollback()
		return err
	}
	roleModel := SystemUserRole{}
	if err := tx.Where("system_user_id=?", u.ID).Delete(&roleModel).Error; err != nil {
		return err
	}
	//如果没有设置权限
	if len(roles) < 1 {
		return tx.Commit().Error
	}
	for _, k := range roles {
		roleModel := SystemRole{Name: k.(string)}
		res := tx.First(&roleModel)
		if res.RowsAffected < 1 || roleModel.State == 0 {
			continue
		}
		userroleModel := SystemUserRole{SystemRoleID: roleModel.ID, SystemUserID: u.ID}
		res = tx.First(&userroleModel)
		if res.RowsAffected > 0 {
			continue
		}
		if err := tx.Create(&userroleModel).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit().Error
}
func (u *SystemUser) UpdatePasswd() error {
	if err := db.Where("id = ?", u.ID).Updates(u).Error; err != nil {
		return err
	}
	return nil
}

func (u *SystemUser) Delete() error {
	if err := db.Model(&u).Where("id = ?", u.ID).Update("state=?", menuStateNo).Error; err != nil {
		return err
	}
	return nil
}
