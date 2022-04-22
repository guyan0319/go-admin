package systemnewdb

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _SystemRoleMenuMgr struct {
	*_BaseMgr
}

// SystemRoleMenuMgr open func
func SystemRoleMenuMgr(db *gorm.DB) *_SystemRoleMenuMgr {
	if db == nil {
		panic(fmt.Errorf("SystemRoleMenuMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SystemRoleMenuMgr{_BaseMgr: &_BaseMgr{DB: db.Table("system_role_menu"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SystemRoleMenuMgr) GetTableName() string {
	return "system_role_menu"
}

// Get 获取
func (obj *_SystemRoleMenuMgr) Get() (result SystemRoleMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemRoleMenu{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SystemRoleMenuMgr) Gets() (results []*SystemRoleMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemRoleMenu{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_SystemRoleMenuMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(SystemRoleMenu{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键
func (obj *_SystemRoleMenuMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithSystemRoleID system_role_id获取 角色主键
func (obj *_SystemRoleMenuMgr) WithSystemRoleID(systemRoleID int) Option {
	return optionFunc(func(o *options) { o.query["system_role_id"] = systemRoleID })
}

// WithSystemMenuID system_menu_id获取 菜单主键
func (obj *_SystemRoleMenuMgr) WithSystemMenuID(systemMenuID int) Option {
	return optionFunc(func(o *options) { o.query["system_menu_id"] = systemMenuID })
}

// GetByOption 功能选项模式获取
func (obj *_SystemRoleMenuMgr) GetByOption(opts ...Option) (result SystemRoleMenu, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SystemRoleMenu{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_SystemRoleMenuMgr) GetByOptions(opts ...Option) (results []*SystemRoleMenu, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SystemRoleMenu{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 主键
func (obj *_SystemRoleMenuMgr) GetFromID(id int) (result SystemRoleMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemRoleMenu{}).Where("`id` = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量查找 主键
func (obj *_SystemRoleMenuMgr) GetBatchFromID(ids []int) (results []*SystemRoleMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemRoleMenu{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromSystemRoleID 通过system_role_id获取内容 角色主键
func (obj *_SystemRoleMenuMgr) GetFromSystemRoleID(systemRoleID int) (results []*SystemRoleMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemRoleMenu{}).Where("`system_role_id` = ?", systemRoleID).Find(&results).Error

	return
}

// GetBatchFromSystemRoleID 批量查找 角色主键
func (obj *_SystemRoleMenuMgr) GetBatchFromSystemRoleID(systemRoleIDs []int) (results []*SystemRoleMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemRoleMenu{}).Where("`system_role_id` IN (?)", systemRoleIDs).Find(&results).Error

	return
}

// GetFromSystemMenuID 通过system_menu_id获取内容 菜单主键
func (obj *_SystemRoleMenuMgr) GetFromSystemMenuID(systemMenuID int) (results []*SystemRoleMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemRoleMenu{}).Where("`system_menu_id` = ?", systemMenuID).Find(&results).Error

	return
}

// GetBatchFromSystemMenuID 批量查找 菜单主键
func (obj *_SystemRoleMenuMgr) GetBatchFromSystemMenuID(systemMenuIDs []int) (results []*SystemRoleMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemRoleMenu{}).Where("`system_menu_id` IN (?)", systemMenuIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_SystemRoleMenuMgr) FetchByPrimaryKey(id int) (result SystemRoleMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemRoleMenu{}).Where("`id` = ?", id).Find(&result).Error

	return
}

// FetchIndexBySystemRoleID  获取多个内容
func (obj *_SystemRoleMenuMgr) FetchIndexBySystemRoleID(systemRoleID int, systemMenuID int) (results []*SystemRoleMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemRoleMenu{}).Where("`system_role_id` = ? AND `system_menu_id` = ?", systemRoleID, systemMenuID).Find(&results).Error

	return
}
