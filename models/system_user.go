package models

import (
	"time"
)

type SystemUser struct {
	Id            int       `json:"id" xorm:"not null pk autoincr comment('主键') INT(11)"`
	Name          string    `json:"name" xorm:"not null comment('姓名') VARCHAR(50)"`
	Nickname      string    `json:"nickname" xorm:"not null default '' comment('用户登录名') unique VARCHAR(50)"`
	Password      string    `json:"password" xorm:"not null comment('密码') index VARCHAR(50)"`
	Salt          string    `json:"salt" xorm:"not null comment('盐') VARCHAR(4)"`
	Phone         string    `json:"phone" xorm:"not null default '' comment('手机号') CHAR(11)"`
	Status        int       `json:"status" xorm:"not null default 1 comment('状态（0 停止1启动）') TINYINT(4)"`
	Utime         time.Time `json:"utime" xorm:"not null default 'CURRENT_TIMESTAMP' comment('更新时间') TIMESTAMP"`
	LastLoginTime time.Time `json:"last_login_time" xorm:"not null default '0000-00-00 00:00:00' comment('上次登录时间') DATETIME"`
	LastLoginIp   string    `json:"last_login_ip" xorm:"not null default '' comment('最近登录IP') VARCHAR(50)"`
	Ctime         time.Time `json:"ctime" xorm:"not null comment('注册时间') DATETIME"`
}
