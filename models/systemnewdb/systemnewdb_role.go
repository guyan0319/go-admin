package systemnewdb

import (
	"strings"
)

var roleStateOk = 1
var roleStateNo = 0

func (r *SystemRole) GetRowMenu() map[int][]string {
	var sr []SystemRole
	res := db.Find(&sr)
	if res.Error != nil {
		panic(res.Error)
	}
	var srMap map[int]string
	srMap = make(map[int]string, 0)
	for _, v := range sr {
		srMap[v.ID] = v.Name
	}
	var srm = SystemRoleMenu{}
	rmArr, _ := srm.GetAll()
	var mrMap = make(map[int][]string, 0)
	for _, value := range rmArr {
		_, ok := srMap[value.SystemRoleID]
		if ok {
			mrMap[value.SystemMenuID] = append(mrMap[value.SystemMenuID], srMap[value.SystemRoleID])
		}
	}
	return mrMap
}
func (r *SystemRole) GetAll() ([]SystemRole, error) {
	var sr []SystemRole
	result := db.Find(&sr)
	return sr, result.Error
}

//根据非零值获取一条数据
func (r *SystemRole) GetRow() (int64, error) {
	result := db.First(&r)
	return result.RowsAffected, result.Error
}
func (r *SystemRole) Delete() error {
	return db.Model(&r).Where("id=?", r.ID).Update("state", roleStateNo).Error
}

func (r *SystemRole) Update(data []int) error {
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
	if err := tx.Model(&r).Where("id = ?", r.ID).Updates(r).Error; err != nil {
		tx.Rollback()
		return err
	}
	if len(data) <= 0 {
		return tx.Commit().Error
	}
	rolemenu := SystemRoleMenu{}
	if err := tx.Where("system_role_id=?", r.ID).Delete(&rolemenu).Error; err != nil {
		return err
	}
	for _, value := range data {
		rm := SystemRoleMenu{SystemRoleID: r.ID, SystemMenuID: value}
		if err := tx.Create(&rm).Error; err != nil {
			return err
		}
	}
	// add Commit() after all actions
	return tx.Commit().Error
}
func (r *SystemRole) AddCommit(data []interface{}) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}
	if err := tx.Create(&r).Error; err != nil {
		return err
	}
	if len(data) <= 0 {
		return tx.Commit().Error
	}

	for _, value := range data {
		menu := SystemMenu{}
		pathMain := value.(map[string]interface{})["path"].(string)
		menu.Path = pathMain
		menu.Component = value.(map[string]interface{})["component"].(string)
		menu.Type = 2
		initMenu := SystemMenu{}
		has := initMenu.GetRowByPathCT(menu)
		if !has {
			continue
		}
		rm := SystemRoleMenu{SystemRoleID: r.ID, SystemMenuID: initMenu.ID}
		if err := tx.Create(&rm).Error; err != nil {
			return err
		}
		children := value.(map[string]interface{})["children"]
		if children == nil {
			continue
		}
		for _, v := range children.([]interface{}) {
			menu := SystemMenu{}
			strings.TrimPrefix(v.(map[string]interface{})["path"].(string), pathMain+"/")
			menu.Component = v.(map[string]interface{})["component"].(string)
			menu.Type = 2
			initMenu := SystemMenu{}
			has := initMenu.GetRowByPathCT(menu)
			if !has {
				continue
			}
			rm := SystemRoleMenu{SystemRoleID: r.ID, SystemMenuID: initMenu.ID}
			if err := tx.Create(&rm).Error; err != nil {
				return err
			}
		}
	}
	return tx.Commit().Error

}

func (r *SystemRole) GetNameList() []string {
	var list []string
	err := db.Model(&r).Where("state=?", 1).Select("name").Find(&list).Error
	if err != nil {
		panic(err)
	}
	return list
}
