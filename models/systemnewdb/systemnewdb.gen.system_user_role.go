package systemnewdb

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _SystemUserRoleMgr struct {
	*_BaseMgr
}

// SystemUserRoleMgr open func
func SystemUserRoleMgr(db *gorm.DB) *_SystemUserRoleMgr {
	if db == nil {
		panic(fmt.Errorf("SystemUserRoleMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SystemUserRoleMgr{_BaseMgr: &_BaseMgr{DB: db.Table("system_user_role"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SystemUserRoleMgr) GetTableName() string {
	return "system_user_role"
}

// Get 获取
func (obj *_SystemUserRoleMgr) Get() (result SystemUserRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemUserRole{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SystemUserRoleMgr) Gets() (results []*SystemUserRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemUserRole{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_SystemUserRoleMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(SystemUserRole{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键
func (obj *_SystemUserRoleMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithSystemUserID system_user_id获取 用户主键
func (obj *_SystemUserRoleMgr) WithSystemUserID(systemUserID int) Option {
	return optionFunc(func(o *options) { o.query["system_user_id"] = systemUserID })
}

// WithSystemRoleID system_role_id获取 角色主键
func (obj *_SystemUserRoleMgr) WithSystemRoleID(systemRoleID int) Option {
	return optionFunc(func(o *options) { o.query["system_role_id"] = systemRoleID })
}

// GetByOption 功能选项模式获取
func (obj *_SystemUserRoleMgr) GetByOption(opts ...Option) (result SystemUserRole, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SystemUserRole{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_SystemUserRoleMgr) GetByOptions(opts ...Option) (results []*SystemUserRole, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SystemUserRole{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 主键
func (obj *_SystemUserRoleMgr) GetFromID(id int) (result SystemUserRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemUserRole{}).Where("`id` = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量查找 主键
func (obj *_SystemUserRoleMgr) GetBatchFromID(ids []int) (results []*SystemUserRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemUserRole{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromSystemUserID 通过system_user_id获取内容 用户主键
func (obj *_SystemUserRoleMgr) GetFromSystemUserID(systemUserID int) (results []*SystemUserRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemUserRole{}).Where("`system_user_id` = ?", systemUserID).Find(&results).Error

	return
}

// GetBatchFromSystemUserID 批量查找 用户主键
func (obj *_SystemUserRoleMgr) GetBatchFromSystemUserID(systemUserIDs []int) (results []*SystemUserRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemUserRole{}).Where("`system_user_id` IN (?)", systemUserIDs).Find(&results).Error

	return
}

// GetFromSystemRoleID 通过system_role_id获取内容 角色主键
func (obj *_SystemUserRoleMgr) GetFromSystemRoleID(systemRoleID int) (results []*SystemUserRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemUserRole{}).Where("`system_role_id` = ?", systemRoleID).Find(&results).Error

	return
}

// GetBatchFromSystemRoleID 批量查找 角色主键
func (obj *_SystemUserRoleMgr) GetBatchFromSystemRoleID(systemRoleIDs []int) (results []*SystemUserRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemUserRole{}).Where("`system_role_id` IN (?)", systemRoleIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_SystemUserRoleMgr) FetchByPrimaryKey(id int) (result SystemUserRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemUserRole{}).Where("`id` = ?", id).Find(&result).Error

	return
}

// FetchIndexBySystemUserID  获取多个内容
func (obj *_SystemUserRoleMgr) FetchIndexBySystemUserID(systemUserID int, systemRoleID int) (results []*SystemUserRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemUserRole{}).Where("`system_user_id` = ? AND `system_role_id` = ?", systemUserID, systemRoleID).Find(&results).Error

	return
}
