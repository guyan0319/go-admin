package systemnewdb

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _SystemUserMgr struct {
	*_BaseMgr
}

// SystemUserMgr open func
func SystemUserMgr(db *gorm.DB) *_SystemUserMgr {
	if db == nil {
		panic(fmt.Errorf("SystemUserMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SystemUserMgr{_BaseMgr: &_BaseMgr{DB: db.Table("system_user"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SystemUserMgr) GetTableName() string {
	return "system_user"
}

// Get 获取
func (obj *_SystemUserMgr) Get() (result SystemUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemUser{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SystemUserMgr) Gets() (results []*SystemUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemUser{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_SystemUserMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(SystemUser{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键
func (obj *_SystemUserMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取 登录名
func (obj *_SystemUserMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithNickname nickname获取 用户昵称
func (obj *_SystemUserMgr) WithNickname(nickname string) Option {
	return optionFunc(func(o *options) { o.query["nickname"] = nickname })
}

// WithPassword password获取 密码
func (obj *_SystemUserMgr) WithPassword(password string) Option {
	return optionFunc(func(o *options) { o.query["password"] = password })
}

// WithSalt salt获取 盐
func (obj *_SystemUserMgr) WithSalt(salt string) Option {
	return optionFunc(func(o *options) { o.query["salt"] = salt })
}

// WithPhone phone获取 手机号
func (obj *_SystemUserMgr) WithPhone(phone string) Option {
	return optionFunc(func(o *options) { o.query["phone"] = phone })
}

// WithAvatar avatar获取 头像
func (obj *_SystemUserMgr) WithAvatar(avatar string) Option {
	return optionFunc(func(o *options) { o.query["avatar"] = avatar })
}

// WithIntroduction introduction获取 简介
func (obj *_SystemUserMgr) WithIntroduction(introduction string) Option {
	return optionFunc(func(o *options) { o.query["introduction"] = introduction })
}

// WithState state获取 状态（0 停止1启动）
func (obj *_SystemUserMgr) WithState(state int8) Option {
	return optionFunc(func(o *options) { o.query["state"] = state })
}

// WithUtime utime获取 更新时间
func (obj *_SystemUserMgr) WithUtime(utime time.Time) Option {
	return optionFunc(func(o *options) { o.query["utime"] = utime })
}

// WithLastLoginTime last_login_time获取 上次登录时间
func (obj *_SystemUserMgr) WithLastLoginTime(lastLoginTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["last_login_time"] = lastLoginTime })
}

// WithLastLoginIP last_login_ip获取 最近登录IP
func (obj *_SystemUserMgr) WithLastLoginIP(lastLoginIP string) Option {
	return optionFunc(func(o *options) { o.query["last_login_ip"] = lastLoginIP })
}

// WithCtime ctime获取 注册时间
func (obj *_SystemUserMgr) WithCtime(ctime time.Time) Option {
	return optionFunc(func(o *options) { o.query["ctime"] = ctime })
}

// GetByOption 功能选项模式获取
func (obj *_SystemUserMgr) GetByOption(opts ...Option) (result SystemUser, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SystemUser{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_SystemUserMgr) GetByOptions(opts ...Option) (results []*SystemUser, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SystemUser{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 主键
func (obj *_SystemUserMgr) GetFromID(id int) (result SystemUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemUser{}).Where("`id` = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量查找 主键
func (obj *_SystemUserMgr) GetBatchFromID(ids []int) (results []*SystemUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemUser{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 登录名
func (obj *_SystemUserMgr) GetFromName(name string) (results []*SystemUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemUser{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 登录名
func (obj *_SystemUserMgr) GetBatchFromName(names []string) (results []*SystemUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemUser{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromNickname 通过nickname获取内容 用户昵称
func (obj *_SystemUserMgr) GetFromNickname(nickname string) (result SystemUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemUser{}).Where("`nickname` = ?", nickname).Find(&result).Error

	return
}

// GetBatchFromNickname 批量查找 用户昵称
func (obj *_SystemUserMgr) GetBatchFromNickname(nicknames []string) (results []*SystemUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemUser{}).Where("`nickname` IN (?)", nicknames).Find(&results).Error

	return
}

// GetFromPassword 通过password获取内容 密码
func (obj *_SystemUserMgr) GetFromPassword(password string) (results []*SystemUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemUser{}).Where("`password` = ?", password).Find(&results).Error

	return
}

// GetBatchFromPassword 批量查找 密码
func (obj *_SystemUserMgr) GetBatchFromPassword(passwords []string) (results []*SystemUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemUser{}).Where("`password` IN (?)", passwords).Find(&results).Error

	return
}

// GetFromSalt 通过salt获取内容 盐
func (obj *_SystemUserMgr) GetFromSalt(salt string) (results []*SystemUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemUser{}).Where("`salt` = ?", salt).Find(&results).Error

	return
}

// GetBatchFromSalt 批量查找 盐
func (obj *_SystemUserMgr) GetBatchFromSalt(salts []string) (results []*SystemUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemUser{}).Where("`salt` IN (?)", salts).Find(&results).Error

	return
}

// GetFromPhone 通过phone获取内容 手机号
func (obj *_SystemUserMgr) GetFromPhone(phone string) (results []*SystemUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemUser{}).Where("`phone` = ?", phone).Find(&results).Error

	return
}

// GetBatchFromPhone 批量查找 手机号
func (obj *_SystemUserMgr) GetBatchFromPhone(phones []string) (results []*SystemUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemUser{}).Where("`phone` IN (?)", phones).Find(&results).Error

	return
}

// GetFromAvatar 通过avatar获取内容 头像
func (obj *_SystemUserMgr) GetFromAvatar(avatar string) (results []*SystemUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemUser{}).Where("`avatar` = ?", avatar).Find(&results).Error

	return
}

// GetBatchFromAvatar 批量查找 头像
func (obj *_SystemUserMgr) GetBatchFromAvatar(avatars []string) (results []*SystemUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemUser{}).Where("`avatar` IN (?)", avatars).Find(&results).Error

	return
}

// GetFromIntroduction 通过introduction获取内容 简介
func (obj *_SystemUserMgr) GetFromIntroduction(introduction string) (results []*SystemUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemUser{}).Where("`introduction` = ?", introduction).Find(&results).Error

	return
}

// GetBatchFromIntroduction 批量查找 简介
func (obj *_SystemUserMgr) GetBatchFromIntroduction(introductions []string) (results []*SystemUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemUser{}).Where("`introduction` IN (?)", introductions).Find(&results).Error

	return
}

// GetFromState 通过state获取内容 状态（0 停止1启动）
func (obj *_SystemUserMgr) GetFromState(state int8) (results []*SystemUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemUser{}).Where("`state` = ?", state).Find(&results).Error

	return
}

// GetBatchFromState 批量查找 状态（0 停止1启动）
func (obj *_SystemUserMgr) GetBatchFromState(states []int8) (results []*SystemUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemUser{}).Where("`state` IN (?)", states).Find(&results).Error

	return
}

// GetFromUtime 通过utime获取内容 更新时间
func (obj *_SystemUserMgr) GetFromUtime(utime time.Time) (results []*SystemUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemUser{}).Where("`utime` = ?", utime).Find(&results).Error

	return
}

// GetBatchFromUtime 批量查找 更新时间
func (obj *_SystemUserMgr) GetBatchFromUtime(utimes []time.Time) (results []*SystemUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemUser{}).Where("`utime` IN (?)", utimes).Find(&results).Error

	return
}

// GetFromLastLoginTime 通过last_login_time获取内容 上次登录时间
func (obj *_SystemUserMgr) GetFromLastLoginTime(lastLoginTime time.Time) (results []*SystemUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemUser{}).Where("`last_login_time` = ?", lastLoginTime).Find(&results).Error

	return
}

// GetBatchFromLastLoginTime 批量查找 上次登录时间
func (obj *_SystemUserMgr) GetBatchFromLastLoginTime(lastLoginTimes []time.Time) (results []*SystemUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemUser{}).Where("`last_login_time` IN (?)", lastLoginTimes).Find(&results).Error

	return
}

// GetFromLastLoginIP 通过last_login_ip获取内容 最近登录IP
func (obj *_SystemUserMgr) GetFromLastLoginIP(lastLoginIP string) (results []*SystemUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemUser{}).Where("`last_login_ip` = ?", lastLoginIP).Find(&results).Error

	return
}

// GetBatchFromLastLoginIP 批量查找 最近登录IP
func (obj *_SystemUserMgr) GetBatchFromLastLoginIP(lastLoginIPs []string) (results []*SystemUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemUser{}).Where("`last_login_ip` IN (?)", lastLoginIPs).Find(&results).Error

	return
}

// GetFromCtime 通过ctime获取内容 注册时间
func (obj *_SystemUserMgr) GetFromCtime(ctime time.Time) (results []*SystemUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemUser{}).Where("`ctime` = ?", ctime).Find(&results).Error

	return
}

// GetBatchFromCtime 批量查找 注册时间
func (obj *_SystemUserMgr) GetBatchFromCtime(ctimes []time.Time) (results []*SystemUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemUser{}).Where("`ctime` IN (?)", ctimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_SystemUserMgr) FetchByPrimaryKey(id int) (result SystemUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemUser{}).Where("`id` = ?", id).Find(&result).Error

	return
}

// FetchUniqueByNICKNAME primary or index 获取唯一内容
func (obj *_SystemUserMgr) FetchUniqueByNICKNAME(nickname string) (result SystemUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemUser{}).Where("`nickname` = ?", nickname).Find(&result).Error

	return
}

// FetchIndexByPASSWORD  获取多个内容
func (obj *_SystemUserMgr) FetchIndexByPASSWORD(password string) (results []*SystemUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemUser{}).Where("`password` = ?", password).Find(&results).Error

	return
}
