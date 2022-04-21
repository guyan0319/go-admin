package systemnewdb

import (
	"time"
)

// SystemArticle [...]
type SystemArticle struct {
	ID              uint64    `gorm:"primaryKey;column:id;type:bigint unsigned;not null" json:"id"`                   // 主键
	Author          int       `gorm:"column:author;type:int;not null;default:0" json:"author"`                        // 作者
	Importance      uint8     `gorm:"column:importance;type:tinyint unsigned;not null;default:1" json:"importance"`   // 重要级别
	State           int8      `gorm:"column:state;type:tinyint;not null;default:0" json:"state"`                      // 状态(0:draft，1:published,10:deleted)
	Title           string    `gorm:"column:title;type:varchar(200);not null;default:''" json:"title"`                // 标题
	Content         string    `gorm:"column:content;type:text;not null" json:"content"`                               // 内容
	ContentShort    string    `gorm:"column:content_short;type:varchar(500);not null;default:''" json:"contentShort"` // 摘要
	SourceURI       string    `gorm:"column:source_uri;type:varchar(200);not null;default:''" json:"sourceUri"`       // 来源
	Ctime           int       `gorm:"column:ctime;type:int;not null;default:0" json:"ctime"`                          // 创建时间
	ImageURI        string    `gorm:"column:image_uri;type:varchar(200);not null;default:''" json:"imageUri"`         // 图片
	CommentDisabled int8      `gorm:"column:comment_disabled;type:tinyint;not null;default:0" json:"commentDisabled"` // 是否展示评论
	DisplayTime     time.Time `gorm:"column:display_time;type:datetime;not null" json:"displayTime"`                  // 发布时间
	Mtime           time.Time `gorm:"column:mtime;type:timestamp;not null" json:"mtime"`                              // 修改时间
}

// SystemArticleColumns get sql column name.获取数据库列名
var SystemArticleColumns = struct {
	ID              string
	Author          string
	Importance      string
	State           string
	Title           string
	Content         string
	ContentShort    string
	SourceURI       string
	Ctime           string
	ImageURI        string
	CommentDisabled string
	DisplayTime     string
	Mtime           string
}{
	ID:              "id",
	Author:          "author",
	Importance:      "importance",
	State:           "state",
	Title:           "title",
	Content:         "content",
	ContentShort:    "content_short",
	SourceURI:       "source_uri",
	Ctime:           "ctime",
	ImageURI:        "image_uri",
	CommentDisabled: "comment_disabled",
	DisplayTime:     "display_time",
	Mtime:           "mtime",
}

// SystemLog 操作日志
type SystemLog struct {
	ID            int64     `gorm:"primaryKey;column:id;type:bigint;not null" json:"id"`                                         // 主键
	SystemUserID  int       `gorm:"index:SYSTEM_USER_ID;column:system_user_id;type:int;default:0" json:"systemUserId"`           // 主键
	Title         string    `gorm:"column:title;type:varchar(300);not null;default:''" json:"title"`                             // 日志标题
	Content       string    `gorm:"column:content;type:text" json:"content"`                                                     // 日志内容记录SQL
	RelationID    int64     `gorm:"index:RELATION_ID;column:relation_id;type:bigint;not null;default:0" json:"relationId"`       // 相关对应表主键
	RelationTable int       `gorm:"index:RELATION_TABLE;column:relation_table;type:int;not null;default:1" json:"relationTable"` // 对应表(1 system_user,2 system_menu,3 system_role)
	IP            string    `gorm:"column:ip;type:varchar(50);not null;default:''" json:"ip"`                                    // ip
	URL           string    `gorm:"column:url;type:varchar(500);not null;default:''" json:"url"`
	Ctime         time.Time `gorm:"index:CTIME;column:ctime;type:datetime;not null;default:0000-00-00 00:00:00" json:"ctime"` // 时间
}

// SystemLogColumns get sql column name.获取数据库列名
var SystemLogColumns = struct {
	ID            string
	SystemUserID  string
	Title         string
	Content       string
	RelationID    string
	RelationTable string
	IP            string
	URL           string
	Ctime         string
}{
	ID:            "id",
	SystemUserID:  "system_user_id",
	Title:         "title",
	Content:       "content",
	RelationID:    "relation_id",
	RelationTable: "relation_table",
	IP:            "ip",
	URL:           "url",
	Ctime:         "ctime",
}

// SystemMenu 权限
type SystemMenu struct {
	ID              int       `gorm:"primaryKey;column:id;type:int;not null" json:"id"`                                    // 主键
	Name            string    `gorm:"column:name;type:varchar(100);not null;default:''" json:"name"`                       // 名称
	Alias           string    `gorm:"column:alias;type:varchar(100);not null;default:''" json:"alias"`                     // 别名
	Path            string    `gorm:"index:path;column:path;type:varchar(50);not null;default:''" json:"path"`             // 路径
	Component       string    `gorm:"column:component;type:varchar(100);not null;default:''" json:"component"`             // 组件
	Redirect        string    `gorm:"column:redirect;type:varchar(200);not null;default:''" json:"redirect"`               // 重定向
	URL             string    `gorm:"column:url;type:varchar(200);not null;default:''" json:"url"`                         // 访问url
	MetaTitle       string    `gorm:"column:meta_title;type:varchar(50);not null;default:''" json:"metaTitle"`             // meta标题
	MetaIcon        string    `gorm:"column:meta_icon;type:varchar(50);not null;default:''" json:"metaIcon"`               // meta icon
	MetaI18n        int8      `gorm:"column:meta_i18n;type:tinyint;not null;default:0" json:"metaI18n"`                    // 是否国家化（1:是 0:否）
	MetaShowlink    int8      `gorm:"column:meta_showlink;type:tinyint;not null;default:0" json:"metaShowlink"`            // 是否总是显示（1:是0：否）
	MetaRank        int8      `gorm:"column:meta_rank;type:tinyint;not null;default:0" json:"metaRank"`                    // 排序
	MetaKeepalive   int8      `gorm:"column:meta_keepalive;type:tinyint;not null;default:0" json:"metaKeepalive"`          // 是否开启缓存（1开启0关闭)
	Type            int8      `gorm:"column:type;type:tinyint;not null;default:2" json:"type"`                             // 类型(1:固定,2:权限配置3特殊)
	MetaFramesrc    string    `gorm:"column:meta_framesrc;type:varchar(300);not null;default:0" json:"metaFramesrc"`       // 内嵌的iframe链接
	TransitionName  int8      `gorm:"column:transition_name;type:tinyint;not null;default:0" json:"transitionName"`        // 是否显示动画（1是0否）
	TransitionEnter string    `gorm:"column:transition_enter;type:varchar(50);not null;default:''" json:"transitionEnter"` // 进入动画
	TransitionLeave string    `gorm:"column:transition_leave;type:varchar(50);not null;default:''" json:"transitionLeave"` // 离开动画
	Dynamiclevel    int8      `gorm:"column:dynamiclevel;type:tinyint;not null;default:3" json:"dynamiclevel"`             // 动态路由可打开的最大数量
	Refreshredirect string    `gorm:"column:refreshredirect;type:varchar(50);not null;default:''" json:"refreshredirect"`  // 刷新重定向
	ExtraiconSvg    int8      `gorm:"column:extraicon_svg;type:tinyint;not null;default:0" json:"extraiconSvg"`            // 额外图标svg(1是0否)
	ExtraiconName   string    `gorm:"column:extraicon_name;type:varchar(50);not null;default:''" json:"extraiconName"`     // 图片名称
	Pid             int       `gorm:"column:pid;type:int;not null;default:0" json:"pid"`                                   // 父ID
	State           int8      `gorm:"column:state;type:tinyint;not null;default:1" json:"state"`                           // 状态（0禁止1启动）
	Level           int8      `gorm:"column:level;type:tinyint;not null;default:0" json:"level"`                           // 层级
	Ctime           time.Time `gorm:"column:ctime;type:datetime;not null" json:"ctime"`                                    // 时间
}

// SystemMenuColumns get sql column name.获取数据库列名
var SystemMenuColumns = struct {
	ID              string
	Name            string
	Alias           string
	Path            string
	Component       string
	Redirect        string
	URL             string
	MetaTitle       string
	MetaIcon        string
	MetaI18n        string
	MetaShowlink    string
	MetaRank        string
	MetaKeepalive   string
	Type            string
	MetaFramesrc    string
	TransitionName  string
	TransitionEnter string
	TransitionLeave string
	Dynamiclevel    string
	Refreshredirect string
	ExtraiconSvg    string
	ExtraiconName   string
	Pid             string
	State           string
	Level           string
	Ctime           string
}{
	ID:              "id",
	Name:            "name",
	Alias:           "alias",
	Path:            "path",
	Component:       "component",
	Redirect:        "redirect",
	URL:             "url",
	MetaTitle:       "meta_title",
	MetaIcon:        "meta_icon",
	MetaI18n:        "meta_i18n",
	MetaShowlink:    "meta_showlink",
	MetaRank:        "meta_rank",
	MetaKeepalive:   "meta_keepalive",
	Type:            "type",
	MetaFramesrc:    "meta_framesrc",
	TransitionName:  "transition_name",
	TransitionEnter: "transition_enter",
	TransitionLeave: "transition_leave",
	Dynamiclevel:    "dynamiclevel",
	Refreshredirect: "refreshredirect",
	ExtraiconSvg:    "extraicon_svg",
	ExtraiconName:   "extraicon_name",
	Pid:             "pid",
	State:           "state",
	Level:           "level",
	Ctime:           "ctime",
}

// SystemRole 角色
type SystemRole struct {
	ID          int       `gorm:"primaryKey;column:id;type:int;not null" json:"id"`                            // 主键
	Name        string    `gorm:"column:name;type:varchar(100);not null" json:"name"`                          // 角色名称
	AliasName   string    `gorm:"column:alias_name;type:varchar(50);not null;default:''" json:"aliasName"`     // 别名
	Description string    `gorm:"column:description;type:varchar(200);not null;default:''" json:"description"` // 描述
	State       int8      `gorm:"index:STATUS;column:state;type:tinyint;not null;default:1" json:"state"`      // 角色状态（0无效1有效）
	Type        int       `gorm:"index:TYPE;column:type;type:int;not null;default:1" json:"type"`              // 属于哪个应用
	Ctime       time.Time `gorm:"column:ctime;type:datetime;not null" json:"ctime"`                            // 创建时间
}

// SystemRoleColumns get sql column name.获取数据库列名
var SystemRoleColumns = struct {
	ID          string
	Name        string
	AliasName   string
	Description string
	State       string
	Type        string
	Ctime       string
}{
	ID:          "id",
	Name:        "name",
	AliasName:   "alias_name",
	Description: "description",
	State:       "state",
	Type:        "type",
	Ctime:       "ctime",
}

// SystemRoleMenu 角色与菜单关联表
type SystemRoleMenu struct {
	ID           int `gorm:"primaryKey;column:id;type:int;not null" json:"id"`                                           // 主键
	SystemRoleID int `gorm:"index:system_role_id;column:system_role_id;type:int;not null;default:0" json:"systemRoleId"` // 角色主键
	SystemMenuID int `gorm:"index:system_role_id;column:system_menu_id;type:int;not null;default:0" json:"systemMenuId"` // 菜单主键
}

// SystemRoleMenuColumns get sql column name.获取数据库列名
var SystemRoleMenuColumns = struct {
	ID           string
	SystemRoleID string
	SystemMenuID string
}{
	ID:           "id",
	SystemRoleID: "system_role_id",
	SystemMenuID: "system_menu_id",
}

// SystemUser 管理账户表
type SystemUser struct {
	ID            int       `gorm:"primaryKey;column:id;type:int;not null" json:"id"`                              // 主键
	Name          string    `gorm:"column:name;type:varchar(50);not null" json:"name"`                             // 登录名
	Nickname      string    `gorm:"unique;column:nickname;type:varchar(50);not null;default:''" json:"nickname"`   // 用户昵称
	Password      string    `gorm:"index:PASSWORD;column:password;type:varchar(50);not null" json:"password"`      // 密码
	Salt          string    `gorm:"column:salt;type:varchar(4);not null" json:"salt"`                              // 盐
	Phone         string    `gorm:"column:phone;type:varchar(11);not null;default:''" json:"phone"`                // 手机号
	Avatar        string    `gorm:"column:avatar;type:varchar(300);not null;default:''" json:"avatar"`             // 头像
	Introduction  string    `gorm:"column:introduction;type:varchar(300);not null;default:''" json:"introduction"` // 简介
	State         int8      `gorm:"column:state;type:tinyint;not null;default:1" json:"state"`                     // 状态（0 停止1启动）
	Utime         time.Time `gorm:"column:utime;type:timestamp;not null" json:"utime"`                             // 更新时间
	LastLoginTime time.Time `gorm:"column:last_login_time;type:datetime;not null" json:"lastLoginTime"`            // 上次登录时间
	LastLoginIP   string    `gorm:"column:last_login_ip;type:varchar(50);not null;default:''" json:"lastLoginIp"`  // 最近登录IP
	Ctime         time.Time `gorm:"column:ctime;type:datetime;not null" json:"ctime"`                              // 注册时间
}

// SystemUserColumns get sql column name.获取数据库列名
var SystemUserColumns = struct {
	ID            string
	Name          string
	Nickname      string
	Password      string
	Salt          string
	Phone         string
	Avatar        string
	Introduction  string
	State         string
	Utime         string
	LastLoginTime string
	LastLoginIP   string
	Ctime         string
}{
	ID:            "id",
	Name:          "name",
	Nickname:      "nickname",
	Password:      "password",
	Salt:          "salt",
	Phone:         "phone",
	Avatar:        "avatar",
	Introduction:  "introduction",
	State:         "state",
	Utime:         "utime",
	LastLoginTime: "last_login_time",
	LastLoginIP:   "last_login_ip",
	Ctime:         "ctime",
}

// SystemUserRole 账户和角色关联表
type SystemUserRole struct {
	ID           int `gorm:"primaryKey;column:id;type:int;not null" json:"id"`                                 // 主键
	SystemUserID int `gorm:"index:system_user_id;column:system_user_id;type:int;not null" json:"systemUserId"` // 用户主键
	SystemRoleID int `gorm:"index:system_user_id;column:system_role_id;type:int;not null" json:"systemRoleId"` // 角色主键
}

// SystemUserRoleColumns get sql column name.获取数据库列名
var SystemUserRoleColumns = struct {
	ID           string
	SystemUserID string
	SystemRoleID string
}{
	ID:           "id",
	SystemUserID: "system_user_id",
	SystemRoleID: "system_role_id",
}
