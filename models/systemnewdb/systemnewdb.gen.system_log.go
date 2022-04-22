package systemnewdb

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _SystemLogMgr struct {
	*_BaseMgr
}

// SystemLogMgr open func
func SystemLogMgr(db *gorm.DB) *_SystemLogMgr {
	if db == nil {
		panic(fmt.Errorf("SystemLogMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SystemLogMgr{_BaseMgr: &_BaseMgr{DB: db.Table("system_log"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SystemLogMgr) GetTableName() string {
	return "system_log"
}

// Get 获取
func (obj *_SystemLogMgr) Get() (result SystemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemLog{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SystemLogMgr) Gets() (results []*SystemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemLog{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_SystemLogMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(SystemLog{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键
func (obj *_SystemLogMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithSystemUserID system_user_id获取 主键
func (obj *_SystemLogMgr) WithSystemUserID(systemUserID int) Option {
	return optionFunc(func(o *options) { o.query["system_user_id"] = systemUserID })
}

// WithTitle title获取 日志标题
func (obj *_SystemLogMgr) WithTitle(title string) Option {
	return optionFunc(func(o *options) { o.query["title"] = title })
}

// WithContent content获取 日志内容记录SQL
func (obj *_SystemLogMgr) WithContent(content string) Option {
	return optionFunc(func(o *options) { o.query["content"] = content })
}

// WithRelationID relation_id获取 相关对应表主键
func (obj *_SystemLogMgr) WithRelationID(relationID int64) Option {
	return optionFunc(func(o *options) { o.query["relation_id"] = relationID })
}

// WithRelationTable relation_table获取 对应表(1 system_user,2 system_menu,3 system_role)
func (obj *_SystemLogMgr) WithRelationTable(relationTable int) Option {
	return optionFunc(func(o *options) { o.query["relation_table"] = relationTable })
}

// WithIP ip获取 ip
func (obj *_SystemLogMgr) WithIP(ip string) Option {
	return optionFunc(func(o *options) { o.query["ip"] = ip })
}

// WithURL url获取
func (obj *_SystemLogMgr) WithURL(url string) Option {
	return optionFunc(func(o *options) { o.query["url"] = url })
}

// WithCtime ctime获取 时间
func (obj *_SystemLogMgr) WithCtime(ctime time.Time) Option {
	return optionFunc(func(o *options) { o.query["ctime"] = ctime })
}

// GetByOption 功能选项模式获取
func (obj *_SystemLogMgr) GetByOption(opts ...Option) (result SystemLog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SystemLog{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_SystemLogMgr) GetByOptions(opts ...Option) (results []*SystemLog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SystemLog{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 主键
func (obj *_SystemLogMgr) GetFromID(id int64) (result SystemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemLog{}).Where("`id` = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量查找 主键
func (obj *_SystemLogMgr) GetBatchFromID(ids []int64) (results []*SystemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemLog{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromSystemUserID 通过system_user_id获取内容 主键
func (obj *_SystemLogMgr) GetFromSystemUserID(systemUserID int) (results []*SystemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemLog{}).Where("`system_user_id` = ?", systemUserID).Find(&results).Error

	return
}

// GetBatchFromSystemUserID 批量查找 主键
func (obj *_SystemLogMgr) GetBatchFromSystemUserID(systemUserIDs []int) (results []*SystemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemLog{}).Where("`system_user_id` IN (?)", systemUserIDs).Find(&results).Error

	return
}

// GetFromTitle 通过title获取内容 日志标题
func (obj *_SystemLogMgr) GetFromTitle(title string) (results []*SystemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemLog{}).Where("`title` = ?", title).Find(&results).Error

	return
}

// GetBatchFromTitle 批量查找 日志标题
func (obj *_SystemLogMgr) GetBatchFromTitle(titles []string) (results []*SystemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemLog{}).Where("`title` IN (?)", titles).Find(&results).Error

	return
}

// GetFromContent 通过content获取内容 日志内容记录SQL
func (obj *_SystemLogMgr) GetFromContent(content string) (results []*SystemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemLog{}).Where("`content` = ?", content).Find(&results).Error

	return
}

// GetBatchFromContent 批量查找 日志内容记录SQL
func (obj *_SystemLogMgr) GetBatchFromContent(contents []string) (results []*SystemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemLog{}).Where("`content` IN (?)", contents).Find(&results).Error

	return
}

// GetFromRelationID 通过relation_id获取内容 相关对应表主键
func (obj *_SystemLogMgr) GetFromRelationID(relationID int64) (results []*SystemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemLog{}).Where("`relation_id` = ?", relationID).Find(&results).Error

	return
}

// GetBatchFromRelationID 批量查找 相关对应表主键
func (obj *_SystemLogMgr) GetBatchFromRelationID(relationIDs []int64) (results []*SystemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemLog{}).Where("`relation_id` IN (?)", relationIDs).Find(&results).Error

	return
}

// GetFromRelationTable 通过relation_table获取内容 对应表(1 system_user,2 system_menu,3 system_role)
func (obj *_SystemLogMgr) GetFromRelationTable(relationTable int) (results []*SystemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemLog{}).Where("`relation_table` = ?", relationTable).Find(&results).Error

	return
}

// GetBatchFromRelationTable 批量查找 对应表(1 system_user,2 system_menu,3 system_role)
func (obj *_SystemLogMgr) GetBatchFromRelationTable(relationTables []int) (results []*SystemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemLog{}).Where("`relation_table` IN (?)", relationTables).Find(&results).Error

	return
}

// GetFromIP 通过ip获取内容 ip
func (obj *_SystemLogMgr) GetFromIP(ip string) (results []*SystemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemLog{}).Where("`ip` = ?", ip).Find(&results).Error

	return
}

// GetBatchFromIP 批量查找 ip
func (obj *_SystemLogMgr) GetBatchFromIP(ips []string) (results []*SystemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemLog{}).Where("`ip` IN (?)", ips).Find(&results).Error

	return
}

// GetFromURL 通过url获取内容
func (obj *_SystemLogMgr) GetFromURL(url string) (results []*SystemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemLog{}).Where("`url` = ?", url).Find(&results).Error

	return
}

// GetBatchFromURL 批量查找
func (obj *_SystemLogMgr) GetBatchFromURL(urls []string) (results []*SystemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemLog{}).Where("`url` IN (?)", urls).Find(&results).Error

	return
}

// GetFromCtime 通过ctime获取内容 时间
func (obj *_SystemLogMgr) GetFromCtime(ctime time.Time) (results []*SystemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemLog{}).Where("`ctime` = ?", ctime).Find(&results).Error

	return
}

// GetBatchFromCtime 批量查找 时间
func (obj *_SystemLogMgr) GetBatchFromCtime(ctimes []time.Time) (results []*SystemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemLog{}).Where("`ctime` IN (?)", ctimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_SystemLogMgr) FetchByPrimaryKey(id int64) (result SystemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemLog{}).Where("`id` = ?", id).Find(&result).Error

	return
}

// FetchIndexBySYSTEMUSERID  获取多个内容
func (obj *_SystemLogMgr) FetchIndexBySYSTEMUSERID(systemUserID int) (results []*SystemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemLog{}).Where("`system_user_id` = ?", systemUserID).Find(&results).Error

	return
}

// FetchIndexByRELATIONID  获取多个内容
func (obj *_SystemLogMgr) FetchIndexByRELATIONID(relationID int64) (results []*SystemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemLog{}).Where("`relation_id` = ?", relationID).Find(&results).Error

	return
}

// FetchIndexByRELATIONTABLE  获取多个内容
func (obj *_SystemLogMgr) FetchIndexByRELATIONTABLE(relationTable int) (results []*SystemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemLog{}).Where("`relation_table` = ?", relationTable).Find(&results).Error

	return
}

// FetchIndexByCTIME  获取多个内容
func (obj *_SystemLogMgr) FetchIndexByCTIME(ctime time.Time) (results []*SystemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemLog{}).Where("`ctime` = ?", ctime).Find(&results).Error

	return
}
