package systemnewdb

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _SystemArticleMgr struct {
	*_BaseMgr
}

// SystemArticleMgr open func
func SystemArticleMgr(db *gorm.DB) *_SystemArticleMgr {
	if db == nil {
		panic(fmt.Errorf("SystemArticleMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SystemArticleMgr{_BaseMgr: &_BaseMgr{DB: db.Table("system_article"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SystemArticleMgr) GetTableName() string {
	return "system_article"
}

// Get 获取
func (obj *_SystemArticleMgr) Get() (result SystemArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemArticle{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SystemArticleMgr) Gets() (results []*SystemArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemArticle{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_SystemArticleMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(SystemArticle{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键
func (obj *_SystemArticleMgr) WithID(id uint64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithAuthor author获取 作者
func (obj *_SystemArticleMgr) WithAuthor(author int) Option {
	return optionFunc(func(o *options) { o.query["author"] = author })
}

// WithImportance importance获取 重要级别
func (obj *_SystemArticleMgr) WithImportance(importance uint8) Option {
	return optionFunc(func(o *options) { o.query["importance"] = importance })
}

// WithState state获取 状态(0:draft，1:published,10:deleted)
func (obj *_SystemArticleMgr) WithState(state int8) Option {
	return optionFunc(func(o *options) { o.query["state"] = state })
}

// WithTitle title获取 标题
func (obj *_SystemArticleMgr) WithTitle(title string) Option {
	return optionFunc(func(o *options) { o.query["title"] = title })
}

// WithContent content获取 内容
func (obj *_SystemArticleMgr) WithContent(content string) Option {
	return optionFunc(func(o *options) { o.query["content"] = content })
}

// WithContentShort content_short获取 摘要
func (obj *_SystemArticleMgr) WithContentShort(contentShort string) Option {
	return optionFunc(func(o *options) { o.query["content_short"] = contentShort })
}

// WithSourceURI source_uri获取 来源
func (obj *_SystemArticleMgr) WithSourceURI(sourceURI string) Option {
	return optionFunc(func(o *options) { o.query["source_uri"] = sourceURI })
}

// WithCtime ctime获取 创建时间
func (obj *_SystemArticleMgr) WithCtime(ctime int) Option {
	return optionFunc(func(o *options) { o.query["ctime"] = ctime })
}

// WithImageURI image_uri获取 图片
func (obj *_SystemArticleMgr) WithImageURI(imageURI string) Option {
	return optionFunc(func(o *options) { o.query["image_uri"] = imageURI })
}

// WithCommentDisabled comment_disabled获取 是否展示评论
func (obj *_SystemArticleMgr) WithCommentDisabled(commentDisabled int8) Option {
	return optionFunc(func(o *options) { o.query["comment_disabled"] = commentDisabled })
}

// WithDisplayTime display_time获取 发布时间
func (obj *_SystemArticleMgr) WithDisplayTime(displayTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["display_time"] = displayTime })
}

// WithMtime mtime获取 修改时间
func (obj *_SystemArticleMgr) WithMtime(mtime time.Time) Option {
	return optionFunc(func(o *options) { o.query["mtime"] = mtime })
}

// GetByOption 功能选项模式获取
func (obj *_SystemArticleMgr) GetByOption(opts ...Option) (result SystemArticle, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SystemArticle{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_SystemArticleMgr) GetByOptions(opts ...Option) (results []*SystemArticle, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SystemArticle{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 主键
func (obj *_SystemArticleMgr) GetFromID(id uint64) (result SystemArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemArticle{}).Where("`id` = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量查找 主键
func (obj *_SystemArticleMgr) GetBatchFromID(ids []uint64) (results []*SystemArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemArticle{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromAuthor 通过author获取内容 作者
func (obj *_SystemArticleMgr) GetFromAuthor(author int) (results []*SystemArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemArticle{}).Where("`author` = ?", author).Find(&results).Error

	return
}

// GetBatchFromAuthor 批量查找 作者
func (obj *_SystemArticleMgr) GetBatchFromAuthor(authors []int) (results []*SystemArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemArticle{}).Where("`author` IN (?)", authors).Find(&results).Error

	return
}

// GetFromImportance 通过importance获取内容 重要级别
func (obj *_SystemArticleMgr) GetFromImportance(importance uint8) (results []*SystemArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemArticle{}).Where("`importance` = ?", importance).Find(&results).Error

	return
}

// GetBatchFromImportance 批量查找 重要级别
func (obj *_SystemArticleMgr) GetBatchFromImportance(importances []uint8) (results []*SystemArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemArticle{}).Where("`importance` IN (?)", importances).Find(&results).Error

	return
}

// GetFromState 通过state获取内容 状态(0:draft，1:published,10:deleted)
func (obj *_SystemArticleMgr) GetFromState(state int8) (results []*SystemArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemArticle{}).Where("`state` = ?", state).Find(&results).Error

	return
}

// GetBatchFromState 批量查找 状态(0:draft，1:published,10:deleted)
func (obj *_SystemArticleMgr) GetBatchFromState(states []int8) (results []*SystemArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemArticle{}).Where("`state` IN (?)", states).Find(&results).Error

	return
}

// GetFromTitle 通过title获取内容 标题
func (obj *_SystemArticleMgr) GetFromTitle(title string) (results []*SystemArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemArticle{}).Where("`title` = ?", title).Find(&results).Error

	return
}

// GetBatchFromTitle 批量查找 标题
func (obj *_SystemArticleMgr) GetBatchFromTitle(titles []string) (results []*SystemArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemArticle{}).Where("`title` IN (?)", titles).Find(&results).Error

	return
}

// GetFromContent 通过content获取内容 内容
func (obj *_SystemArticleMgr) GetFromContent(content string) (results []*SystemArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemArticle{}).Where("`content` = ?", content).Find(&results).Error

	return
}

// GetBatchFromContent 批量查找 内容
func (obj *_SystemArticleMgr) GetBatchFromContent(contents []string) (results []*SystemArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemArticle{}).Where("`content` IN (?)", contents).Find(&results).Error

	return
}

// GetFromContentShort 通过content_short获取内容 摘要
func (obj *_SystemArticleMgr) GetFromContentShort(contentShort string) (results []*SystemArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemArticle{}).Where("`content_short` = ?", contentShort).Find(&results).Error

	return
}

// GetBatchFromContentShort 批量查找 摘要
func (obj *_SystemArticleMgr) GetBatchFromContentShort(contentShorts []string) (results []*SystemArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemArticle{}).Where("`content_short` IN (?)", contentShorts).Find(&results).Error

	return
}

// GetFromSourceURI 通过source_uri获取内容 来源
func (obj *_SystemArticleMgr) GetFromSourceURI(sourceURI string) (results []*SystemArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemArticle{}).Where("`source_uri` = ?", sourceURI).Find(&results).Error

	return
}

// GetBatchFromSourceURI 批量查找 来源
func (obj *_SystemArticleMgr) GetBatchFromSourceURI(sourceURIs []string) (results []*SystemArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemArticle{}).Where("`source_uri` IN (?)", sourceURIs).Find(&results).Error

	return
}

// GetFromCtime 通过ctime获取内容 创建时间
func (obj *_SystemArticleMgr) GetFromCtime(ctime int) (results []*SystemArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemArticle{}).Where("`ctime` = ?", ctime).Find(&results).Error

	return
}

// GetBatchFromCtime 批量查找 创建时间
func (obj *_SystemArticleMgr) GetBatchFromCtime(ctimes []int) (results []*SystemArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemArticle{}).Where("`ctime` IN (?)", ctimes).Find(&results).Error

	return
}

// GetFromImageURI 通过image_uri获取内容 图片
func (obj *_SystemArticleMgr) GetFromImageURI(imageURI string) (results []*SystemArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemArticle{}).Where("`image_uri` = ?", imageURI).Find(&results).Error

	return
}

// GetBatchFromImageURI 批量查找 图片
func (obj *_SystemArticleMgr) GetBatchFromImageURI(imageURIs []string) (results []*SystemArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemArticle{}).Where("`image_uri` IN (?)", imageURIs).Find(&results).Error

	return
}

// GetFromCommentDisabled 通过comment_disabled获取内容 是否展示评论
func (obj *_SystemArticleMgr) GetFromCommentDisabled(commentDisabled int8) (results []*SystemArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemArticle{}).Where("`comment_disabled` = ?", commentDisabled).Find(&results).Error

	return
}

// GetBatchFromCommentDisabled 批量查找 是否展示评论
func (obj *_SystemArticleMgr) GetBatchFromCommentDisabled(commentDisableds []int8) (results []*SystemArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemArticle{}).Where("`comment_disabled` IN (?)", commentDisableds).Find(&results).Error

	return
}

// GetFromDisplayTime 通过display_time获取内容 发布时间
func (obj *_SystemArticleMgr) GetFromDisplayTime(displayTime time.Time) (results []*SystemArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemArticle{}).Where("`display_time` = ?", displayTime).Find(&results).Error

	return
}

// GetBatchFromDisplayTime 批量查找 发布时间
func (obj *_SystemArticleMgr) GetBatchFromDisplayTime(displayTimes []time.Time) (results []*SystemArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemArticle{}).Where("`display_time` IN (?)", displayTimes).Find(&results).Error

	return
}

// GetFromMtime 通过mtime获取内容 修改时间
func (obj *_SystemArticleMgr) GetFromMtime(mtime time.Time) (results []*SystemArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemArticle{}).Where("`mtime` = ?", mtime).Find(&results).Error

	return
}

// GetBatchFromMtime 批量查找 修改时间
func (obj *_SystemArticleMgr) GetBatchFromMtime(mtimes []time.Time) (results []*SystemArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemArticle{}).Where("`mtime` IN (?)", mtimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_SystemArticleMgr) FetchByPrimaryKey(id uint64) (result SystemArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemArticle{}).Where("`id` = ?", id).Find(&result).Error

	return
}
