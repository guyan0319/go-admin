package systemnewdb

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _SystemMenuMgr struct {
	*_BaseMgr
}

// SystemMenuMgr open func
func SystemMenuMgr(db *gorm.DB) *_SystemMenuMgr {
	if db == nil {
		panic(fmt.Errorf("SystemMenuMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SystemMenuMgr{_BaseMgr: &_BaseMgr{DB: db.Table("system_menu"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SystemMenuMgr) GetTableName() string {
	return "system_menu"
}

// Get 获取
func (obj *_SystemMenuMgr) Get() (result SystemMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemMenu{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SystemMenuMgr) Gets() (results []*SystemMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemMenu{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_SystemMenuMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(SystemMenu{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键
func (obj *_SystemMenuMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取 名称
func (obj *_SystemMenuMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithAlias alias获取 别名
func (obj *_SystemMenuMgr) WithAlias(alias string) Option {
	return optionFunc(func(o *options) { o.query["alias"] = alias })
}

// WithPath path获取 路径
func (obj *_SystemMenuMgr) WithPath(path string) Option {
	return optionFunc(func(o *options) { o.query["path"] = path })
}

// WithComponent component获取 组件
func (obj *_SystemMenuMgr) WithComponent(component string) Option {
	return optionFunc(func(o *options) { o.query["component"] = component })
}

// WithRedirect redirect获取 重定向
func (obj *_SystemMenuMgr) WithRedirect(redirect string) Option {
	return optionFunc(func(o *options) { o.query["redirect"] = redirect })
}

// WithURL url获取 访问url
func (obj *_SystemMenuMgr) WithURL(url string) Option {
	return optionFunc(func(o *options) { o.query["url"] = url })
}

// WithMetaTitle meta_title获取 meta标题
func (obj *_SystemMenuMgr) WithMetaTitle(metaTitle string) Option {
	return optionFunc(func(o *options) { o.query["meta_title"] = metaTitle })
}

// WithMetaIcon meta_icon获取 meta icon
func (obj *_SystemMenuMgr) WithMetaIcon(metaIcon string) Option {
	return optionFunc(func(o *options) { o.query["meta_icon"] = metaIcon })
}

// WithMetaI18n meta_i18n获取 是否国家化（1:是 0:否）
func (obj *_SystemMenuMgr) WithMetaI18n(metaI18n int8) Option {
	return optionFunc(func(o *options) { o.query["meta_i18n"] = metaI18n })
}

// WithMetaShowlink meta_showlink获取 是否总是显示（1:是0：否）
func (obj *_SystemMenuMgr) WithMetaShowlink(metaShowlink int8) Option {
	return optionFunc(func(o *options) { o.query["meta_showlink"] = metaShowlink })
}

// WithMetaShowparent meta_showparent获取 是否显示父级菜单1是0否
func (obj *_SystemMenuMgr) WithMetaShowparent(metaShowparent int8) Option {
	return optionFunc(func(o *options) { o.query["meta_showparent"] = metaShowparent })
}

// WithMetaRank meta_rank获取 排序
func (obj *_SystemMenuMgr) WithMetaRank(metaRank int8) Option {
	return optionFunc(func(o *options) { o.query["meta_rank"] = metaRank })
}

// WithMetaKeepalive meta_keepalive获取 是否开启缓存（1开启0关闭)
func (obj *_SystemMenuMgr) WithMetaKeepalive(metaKeepalive int8) Option {
	return optionFunc(func(o *options) { o.query["meta_keepalive"] = metaKeepalive })
}

// WithType type获取 类型(1:固定,2:权限配置3特殊)
func (obj *_SystemMenuMgr) WithType(_type int8) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithMetaFramesrc meta_framesrc获取 内嵌的iframe链接
func (obj *_SystemMenuMgr) WithMetaFramesrc(metaFramesrc string) Option {
	return optionFunc(func(o *options) { o.query["meta_framesrc"] = metaFramesrc })
}

// WithTransitionName transition_name获取 当前路由动画效果
//当前路由动画效果
func (obj *_SystemMenuMgr) WithTransitionName(transitionName string) Option {
	return optionFunc(func(o *options) { o.query["transition_name"] = transitionName })
}

// WithTransitionEnter transition_enter获取 进入动画
func (obj *_SystemMenuMgr) WithTransitionEnter(transitionEnter string) Option {
	return optionFunc(func(o *options) { o.query["transition_enter"] = transitionEnter })
}

// WithTransitionLeave transition_leave获取 离开动画
func (obj *_SystemMenuMgr) WithTransitionLeave(transitionLeave string) Option {
	return optionFunc(func(o *options) { o.query["transition_leave"] = transitionLeave })
}

// WithDynamiclevel dynamiclevel获取 动态路由可打开的最大数量
func (obj *_SystemMenuMgr) WithDynamiclevel(dynamiclevel int8) Option {
	return optionFunc(func(o *options) { o.query["dynamiclevel"] = dynamiclevel })
}

// WithRefreshredirect refreshredirect获取 刷新重定向
func (obj *_SystemMenuMgr) WithRefreshredirect(refreshredirect string) Option {
	return optionFunc(func(o *options) { o.query["refreshredirect"] = refreshredirect })
}

// WithExtraiconSvg extraicon_svg获取 额外图标svg(1是0否)
func (obj *_SystemMenuMgr) WithExtraiconSvg(extraiconSvg int8) Option {
	return optionFunc(func(o *options) { o.query["extraicon_svg"] = extraiconSvg })
}

// WithExtraiconName extraicon_name获取 图片名称
func (obj *_SystemMenuMgr) WithExtraiconName(extraiconName string) Option {
	return optionFunc(func(o *options) { o.query["extraicon_name"] = extraiconName })
}

// WithPid pid获取 父ID
func (obj *_SystemMenuMgr) WithPid(pid int) Option {
	return optionFunc(func(o *options) { o.query["pid"] = pid })
}

// WithState state获取 状态（0禁止1启动）
func (obj *_SystemMenuMgr) WithState(state int8) Option {
	return optionFunc(func(o *options) { o.query["state"] = state })
}

// WithLevel level获取 层级
func (obj *_SystemMenuMgr) WithLevel(level int8) Option {
	return optionFunc(func(o *options) { o.query["level"] = level })
}

// WithCtime ctime获取 时间
func (obj *_SystemMenuMgr) WithCtime(ctime time.Time) Option {
	return optionFunc(func(o *options) { o.query["ctime"] = ctime })
}

// GetByOption 功能选项模式获取
func (obj *_SystemMenuMgr) GetByOption(opts ...Option) (result SystemMenu, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SystemMenu{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_SystemMenuMgr) GetByOptions(opts ...Option) (results []*SystemMenu, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SystemMenu{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 主键
func (obj *_SystemMenuMgr) GetFromID(id int) (result SystemMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemMenu{}).Where("`id` = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量查找 主键
func (obj *_SystemMenuMgr) GetBatchFromID(ids []int) (results []*SystemMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemMenu{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 名称
func (obj *_SystemMenuMgr) GetFromName(name string) (results []*SystemMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemMenu{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 名称
func (obj *_SystemMenuMgr) GetBatchFromName(names []string) (results []*SystemMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemMenu{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromAlias 通过alias获取内容 别名
func (obj *_SystemMenuMgr) GetFromAlias(alias string) (results []*SystemMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemMenu{}).Where("`alias` = ?", alias).Find(&results).Error

	return
}

// GetBatchFromAlias 批量查找 别名
func (obj *_SystemMenuMgr) GetBatchFromAlias(aliass []string) (results []*SystemMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemMenu{}).Where("`alias` IN (?)", aliass).Find(&results).Error

	return
}

// GetFromPath 通过path获取内容 路径
func (obj *_SystemMenuMgr) GetFromPath(path string) (results []*SystemMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemMenu{}).Where("`path` = ?", path).Find(&results).Error

	return
}

// GetBatchFromPath 批量查找 路径
func (obj *_SystemMenuMgr) GetBatchFromPath(paths []string) (results []*SystemMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemMenu{}).Where("`path` IN (?)", paths).Find(&results).Error

	return
}

// GetFromComponent 通过component获取内容 组件
func (obj *_SystemMenuMgr) GetFromComponent(component string) (results []*SystemMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemMenu{}).Where("`component` = ?", component).Find(&results).Error

	return
}

// GetBatchFromComponent 批量查找 组件
func (obj *_SystemMenuMgr) GetBatchFromComponent(components []string) (results []*SystemMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemMenu{}).Where("`component` IN (?)", components).Find(&results).Error

	return
}

// GetFromRedirect 通过redirect获取内容 重定向
func (obj *_SystemMenuMgr) GetFromRedirect(redirect string) (results []*SystemMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemMenu{}).Where("`redirect` = ?", redirect).Find(&results).Error

	return
}

// GetBatchFromRedirect 批量查找 重定向
func (obj *_SystemMenuMgr) GetBatchFromRedirect(redirects []string) (results []*SystemMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemMenu{}).Where("`redirect` IN (?)", redirects).Find(&results).Error

	return
}

// GetFromURL 通过url获取内容 访问url
func (obj *_SystemMenuMgr) GetFromURL(url string) (results []*SystemMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemMenu{}).Where("`url` = ?", url).Find(&results).Error

	return
}

// GetBatchFromURL 批量查找 访问url
func (obj *_SystemMenuMgr) GetBatchFromURL(urls []string) (results []*SystemMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemMenu{}).Where("`url` IN (?)", urls).Find(&results).Error

	return
}

// GetFromMetaTitle 通过meta_title获取内容 meta标题
func (obj *_SystemMenuMgr) GetFromMetaTitle(metaTitle string) (results []*SystemMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemMenu{}).Where("`meta_title` = ?", metaTitle).Find(&results).Error

	return
}

// GetBatchFromMetaTitle 批量查找 meta标题
func (obj *_SystemMenuMgr) GetBatchFromMetaTitle(metaTitles []string) (results []*SystemMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemMenu{}).Where("`meta_title` IN (?)", metaTitles).Find(&results).Error

	return
}

// GetFromMetaIcon 通过meta_icon获取内容 meta icon
func (obj *_SystemMenuMgr) GetFromMetaIcon(metaIcon string) (results []*SystemMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemMenu{}).Where("`meta_icon` = ?", metaIcon).Find(&results).Error

	return
}

// GetBatchFromMetaIcon 批量查找 meta icon
func (obj *_SystemMenuMgr) GetBatchFromMetaIcon(metaIcons []string) (results []*SystemMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemMenu{}).Where("`meta_icon` IN (?)", metaIcons).Find(&results).Error

	return
}

// GetFromMetaI18n 通过meta_i18n获取内容 是否国家化（1:是 0:否）
func (obj *_SystemMenuMgr) GetFromMetaI18n(metaI18n int8) (results []*SystemMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemMenu{}).Where("`meta_i18n` = ?", metaI18n).Find(&results).Error

	return
}

// GetBatchFromMetaI18n 批量查找 是否国家化（1:是 0:否）
func (obj *_SystemMenuMgr) GetBatchFromMetaI18n(metaI18ns []int8) (results []*SystemMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemMenu{}).Where("`meta_i18n` IN (?)", metaI18ns).Find(&results).Error

	return
}

// GetFromMetaShowlink 通过meta_showlink获取内容 是否总是显示（1:是0：否）
func (obj *_SystemMenuMgr) GetFromMetaShowlink(metaShowlink int8) (results []*SystemMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemMenu{}).Where("`meta_showlink` = ?", metaShowlink).Find(&results).Error

	return
}

// GetBatchFromMetaShowlink 批量查找 是否总是显示（1:是0：否）
func (obj *_SystemMenuMgr) GetBatchFromMetaShowlink(metaShowlinks []int8) (results []*SystemMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemMenu{}).Where("`meta_showlink` IN (?)", metaShowlinks).Find(&results).Error

	return
}

// GetFromMetaShowparent 通过meta_showparent获取内容 是否显示父级菜单1是0否
func (obj *_SystemMenuMgr) GetFromMetaShowparent(metaShowparent int8) (results []*SystemMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemMenu{}).Where("`meta_showparent` = ?", metaShowparent).Find(&results).Error

	return
}

// GetBatchFromMetaShowparent 批量查找 是否显示父级菜单1是0否
func (obj *_SystemMenuMgr) GetBatchFromMetaShowparent(metaShowparents []int8) (results []*SystemMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemMenu{}).Where("`meta_showparent` IN (?)", metaShowparents).Find(&results).Error

	return
}

// GetFromMetaRank 通过meta_rank获取内容 排序
func (obj *_SystemMenuMgr) GetFromMetaRank(metaRank int8) (results []*SystemMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemMenu{}).Where("`meta_rank` = ?", metaRank).Find(&results).Error

	return
}

// GetBatchFromMetaRank 批量查找 排序
func (obj *_SystemMenuMgr) GetBatchFromMetaRank(metaRanks []int8) (results []*SystemMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemMenu{}).Where("`meta_rank` IN (?)", metaRanks).Find(&results).Error

	return
}

// GetFromMetaKeepalive 通过meta_keepalive获取内容 是否开启缓存（1开启0关闭)
func (obj *_SystemMenuMgr) GetFromMetaKeepalive(metaKeepalive int8) (results []*SystemMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemMenu{}).Where("`meta_keepalive` = ?", metaKeepalive).Find(&results).Error

	return
}

// GetBatchFromMetaKeepalive 批量查找 是否开启缓存（1开启0关闭)
func (obj *_SystemMenuMgr) GetBatchFromMetaKeepalive(metaKeepalives []int8) (results []*SystemMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemMenu{}).Where("`meta_keepalive` IN (?)", metaKeepalives).Find(&results).Error

	return
}

// GetFromType 通过type获取内容 类型(1:固定,2:权限配置3特殊)
func (obj *_SystemMenuMgr) GetFromType(_type int8) (results []*SystemMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemMenu{}).Where("`type` = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量查找 类型(1:固定,2:权限配置3特殊)
func (obj *_SystemMenuMgr) GetBatchFromType(_types []int8) (results []*SystemMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemMenu{}).Where("`type` IN (?)", _types).Find(&results).Error

	return
}

// GetFromMetaFramesrc 通过meta_framesrc获取内容 内嵌的iframe链接
func (obj *_SystemMenuMgr) GetFromMetaFramesrc(metaFramesrc string) (results []*SystemMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemMenu{}).Where("`meta_framesrc` = ?", metaFramesrc).Find(&results).Error

	return
}

// GetBatchFromMetaFramesrc 批量查找 内嵌的iframe链接
func (obj *_SystemMenuMgr) GetBatchFromMetaFramesrc(metaFramesrcs []string) (results []*SystemMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemMenu{}).Where("`meta_framesrc` IN (?)", metaFramesrcs).Find(&results).Error

	return
}

// GetFromTransitionName 通过transition_name获取内容 当前路由动画效果
//当前路由动画效果
func (obj *_SystemMenuMgr) GetFromTransitionName(transitionName string) (results []*SystemMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemMenu{}).Where("`transition_name` = ?", transitionName).Find(&results).Error

	return
}

// GetBatchFromTransitionName 批量查找 当前路由动画效果
//当前路由动画效果
func (obj *_SystemMenuMgr) GetBatchFromTransitionName(transitionNames []string) (results []*SystemMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemMenu{}).Where("`transition_name` IN (?)", transitionNames).Find(&results).Error

	return
}

// GetFromTransitionEnter 通过transition_enter获取内容 进入动画
func (obj *_SystemMenuMgr) GetFromTransitionEnter(transitionEnter string) (results []*SystemMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemMenu{}).Where("`transition_enter` = ?", transitionEnter).Find(&results).Error

	return
}

// GetBatchFromTransitionEnter 批量查找 进入动画
func (obj *_SystemMenuMgr) GetBatchFromTransitionEnter(transitionEnters []string) (results []*SystemMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemMenu{}).Where("`transition_enter` IN (?)", transitionEnters).Find(&results).Error

	return
}

// GetFromTransitionLeave 通过transition_leave获取内容 离开动画
func (obj *_SystemMenuMgr) GetFromTransitionLeave(transitionLeave string) (results []*SystemMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemMenu{}).Where("`transition_leave` = ?", transitionLeave).Find(&results).Error

	return
}

// GetBatchFromTransitionLeave 批量查找 离开动画
func (obj *_SystemMenuMgr) GetBatchFromTransitionLeave(transitionLeaves []string) (results []*SystemMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemMenu{}).Where("`transition_leave` IN (?)", transitionLeaves).Find(&results).Error

	return
}

// GetFromDynamiclevel 通过dynamiclevel获取内容 动态路由可打开的最大数量
func (obj *_SystemMenuMgr) GetFromDynamiclevel(dynamiclevel int8) (results []*SystemMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemMenu{}).Where("`dynamiclevel` = ?", dynamiclevel).Find(&results).Error

	return
}

// GetBatchFromDynamiclevel 批量查找 动态路由可打开的最大数量
func (obj *_SystemMenuMgr) GetBatchFromDynamiclevel(dynamiclevels []int8) (results []*SystemMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemMenu{}).Where("`dynamiclevel` IN (?)", dynamiclevels).Find(&results).Error

	return
}

// GetFromRefreshredirect 通过refreshredirect获取内容 刷新重定向
func (obj *_SystemMenuMgr) GetFromRefreshredirect(refreshredirect string) (results []*SystemMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemMenu{}).Where("`refreshredirect` = ?", refreshredirect).Find(&results).Error

	return
}

// GetBatchFromRefreshredirect 批量查找 刷新重定向
func (obj *_SystemMenuMgr) GetBatchFromRefreshredirect(refreshredirects []string) (results []*SystemMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemMenu{}).Where("`refreshredirect` IN (?)", refreshredirects).Find(&results).Error

	return
}

// GetFromExtraiconSvg 通过extraicon_svg获取内容 额外图标svg(1是0否)
func (obj *_SystemMenuMgr) GetFromExtraiconSvg(extraiconSvg int8) (results []*SystemMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemMenu{}).Where("`extraicon_svg` = ?", extraiconSvg).Find(&results).Error

	return
}

// GetBatchFromExtraiconSvg 批量查找 额外图标svg(1是0否)
func (obj *_SystemMenuMgr) GetBatchFromExtraiconSvg(extraiconSvgs []int8) (results []*SystemMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemMenu{}).Where("`extraicon_svg` IN (?)", extraiconSvgs).Find(&results).Error

	return
}

// GetFromExtraiconName 通过extraicon_name获取内容 图片名称
func (obj *_SystemMenuMgr) GetFromExtraiconName(extraiconName string) (results []*SystemMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemMenu{}).Where("`extraicon_name` = ?", extraiconName).Find(&results).Error

	return
}

// GetBatchFromExtraiconName 批量查找 图片名称
func (obj *_SystemMenuMgr) GetBatchFromExtraiconName(extraiconNames []string) (results []*SystemMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemMenu{}).Where("`extraicon_name` IN (?)", extraiconNames).Find(&results).Error

	return
}

// GetFromPid 通过pid获取内容 父ID
func (obj *_SystemMenuMgr) GetFromPid(pid int) (results []*SystemMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemMenu{}).Where("`pid` = ?", pid).Find(&results).Error

	return
}

// GetBatchFromPid 批量查找 父ID
func (obj *_SystemMenuMgr) GetBatchFromPid(pids []int) (results []*SystemMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemMenu{}).Where("`pid` IN (?)", pids).Find(&results).Error

	return
}

// GetFromState 通过state获取内容 状态（0禁止1启动）
func (obj *_SystemMenuMgr) GetFromState(state int8) (results []*SystemMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemMenu{}).Where("`state` = ?", state).Find(&results).Error

	return
}

// GetBatchFromState 批量查找 状态（0禁止1启动）
func (obj *_SystemMenuMgr) GetBatchFromState(states []int8) (results []*SystemMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemMenu{}).Where("`state` IN (?)", states).Find(&results).Error

	return
}

// GetFromLevel 通过level获取内容 层级
func (obj *_SystemMenuMgr) GetFromLevel(level int8) (results []*SystemMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemMenu{}).Where("`level` = ?", level).Find(&results).Error

	return
}

// GetBatchFromLevel 批量查找 层级
func (obj *_SystemMenuMgr) GetBatchFromLevel(levels []int8) (results []*SystemMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemMenu{}).Where("`level` IN (?)", levels).Find(&results).Error

	return
}

// GetFromCtime 通过ctime获取内容 时间
func (obj *_SystemMenuMgr) GetFromCtime(ctime time.Time) (results []*SystemMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemMenu{}).Where("`ctime` = ?", ctime).Find(&results).Error

	return
}

// GetBatchFromCtime 批量查找 时间
func (obj *_SystemMenuMgr) GetBatchFromCtime(ctimes []time.Time) (results []*SystemMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemMenu{}).Where("`ctime` IN (?)", ctimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_SystemMenuMgr) FetchByPrimaryKey(id int) (result SystemMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemMenu{}).Where("`id` = ?", id).Find(&result).Error

	return
}

// FetchIndexByPath  获取多个内容
func (obj *_SystemMenuMgr) FetchIndexByPath(path string) (results []*SystemMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemMenu{}).Where("`path` = ?", path).Find(&results).Error

	return
}
