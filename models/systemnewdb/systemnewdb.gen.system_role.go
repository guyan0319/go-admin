package systemnewdb

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _SystemRoleMgr struct {
	*_BaseMgr
}

// SystemRoleMgr open func
func SystemRoleMgr(db *gorm.DB) *_SystemRoleMgr {
	if db == nil {
		panic(fmt.Errorf("SystemRoleMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SystemRoleMgr{_BaseMgr: &_BaseMgr{DB: db.Table("system_role"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SystemRoleMgr) GetTableName() string {
	return "system_role"
}

// Get 获取
func (obj *_SystemRoleMgr) Get() (result SystemRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemRole{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SystemRoleMgr) Gets() (results []*SystemRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemRole{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_SystemRoleMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(SystemRole{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键
func (obj *_SystemRoleMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取 角色名称
func (obj *_SystemRoleMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithAliasName alias_name获取 别名
func (obj *_SystemRoleMgr) WithAliasName(aliasName string) Option {
	return optionFunc(func(o *options) { o.query["alias_name"] = aliasName })
}

// WithDescription description获取 描述
func (obj *_SystemRoleMgr) WithDescription(description string) Option {
	return optionFunc(func(o *options) { o.query["description"] = description })
}

// WithState state获取 角色状态（0无效1有效）
func (obj *_SystemRoleMgr) WithState(state int8) Option {
	return optionFunc(func(o *options) { o.query["state"] = state })
}

// WithType type获取 属于哪个应用
func (obj *_SystemRoleMgr) WithType(_type int) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithCtime ctime获取 创建时间
func (obj *_SystemRoleMgr) WithCtime(ctime time.Time) Option {
	return optionFunc(func(o *options) { o.query["ctime"] = ctime })
}

// GetByOption 功能选项模式获取
func (obj *_SystemRoleMgr) GetByOption(opts ...Option) (result SystemRole, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SystemRole{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_SystemRoleMgr) GetByOptions(opts ...Option) (results []*SystemRole, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SystemRole{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 主键
func (obj *_SystemRoleMgr) GetFromID(id int) (result SystemRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemRole{}).Where("`id` = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量查找 主键
func (obj *_SystemRoleMgr) GetBatchFromID(ids []int) (results []*SystemRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemRole{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 角色名称
func (obj *_SystemRoleMgr) GetFromName(name string) (results []*SystemRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemRole{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 角色名称
func (obj *_SystemRoleMgr) GetBatchFromName(names []string) (results []*SystemRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemRole{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromAliasName 通过alias_name获取内容 别名
func (obj *_SystemRoleMgr) GetFromAliasName(aliasName string) (results []*SystemRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemRole{}).Where("`alias_name` = ?", aliasName).Find(&results).Error

	return
}

// GetBatchFromAliasName 批量查找 别名
func (obj *_SystemRoleMgr) GetBatchFromAliasName(aliasNames []string) (results []*SystemRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemRole{}).Where("`alias_name` IN (?)", aliasNames).Find(&results).Error

	return
}

// GetFromDescription 通过description获取内容 描述
func (obj *_SystemRoleMgr) GetFromDescription(description string) (results []*SystemRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemRole{}).Where("`description` = ?", description).Find(&results).Error

	return
}

// GetBatchFromDescription 批量查找 描述
func (obj *_SystemRoleMgr) GetBatchFromDescription(descriptions []string) (results []*SystemRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemRole{}).Where("`description` IN (?)", descriptions).Find(&results).Error

	return
}

// GetFromState 通过state获取内容 角色状态（0无效1有效）
func (obj *_SystemRoleMgr) GetFromState(state int8) (results []*SystemRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemRole{}).Where("`state` = ?", state).Find(&results).Error

	return
}

// GetBatchFromState 批量查找 角色状态（0无效1有效）
func (obj *_SystemRoleMgr) GetBatchFromState(states []int8) (results []*SystemRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemRole{}).Where("`state` IN (?)", states).Find(&results).Error

	return
}

// GetFromType 通过type获取内容 属于哪个应用
func (obj *_SystemRoleMgr) GetFromType(_type int) (results []*SystemRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemRole{}).Where("`type` = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量查找 属于哪个应用
func (obj *_SystemRoleMgr) GetBatchFromType(_types []int) (results []*SystemRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemRole{}).Where("`type` IN (?)", _types).Find(&results).Error

	return
}

// GetFromCtime 通过ctime获取内容 创建时间
func (obj *_SystemRoleMgr) GetFromCtime(ctime time.Time) (results []*SystemRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemRole{}).Where("`ctime` = ?", ctime).Find(&results).Error

	return
}

// GetBatchFromCtime 批量查找 创建时间
func (obj *_SystemRoleMgr) GetBatchFromCtime(ctimes []time.Time) (results []*SystemRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemRole{}).Where("`ctime` IN (?)", ctimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_SystemRoleMgr) FetchByPrimaryKey(id int) (result SystemRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemRole{}).Where("`id` = ?", id).Find(&result).Error

	return
}

// FetchIndexBySTATUS  获取多个内容
func (obj *_SystemRoleMgr) FetchIndexBySTATUS(state int8) (results []*SystemRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemRole{}).Where("`state` = ?", state).Find(&results).Error

	return
}

// FetchIndexByTYPE  获取多个内容
func (obj *_SystemRoleMgr) FetchIndexByTYPE(_type int) (results []*SystemRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemRole{}).Where("`type` = ?", _type).Find(&results).Error

	return
}
