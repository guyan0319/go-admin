package models

import (
	"go-admin/public/common"
	"time"
)

type SystemUser struct {
	Id            int       `json:"id" xorm:"not null pk autoincr comment('主键') INT(11)"`
	Name          string    `json:"name" xorm:"not null comment('姓名') VARCHAR(50)"`
	Nickname      string    `json:"nickname" xorm:"not null default '' comment('用户登录名') unique VARCHAR(50)"`
	Password      string    `json:"password" xorm:"not null comment('密码') index VARCHAR(50)"`
	Salt          string    `json:"salt" xorm:"not null comment('盐') VARCHAR(4)"`
	Phone         string    `json:"phone" xorm:"not null default '' comment('手机号') VARCHAR(11)"`
	Avatar        string    `json:"avatar" xorm:"not null default '' comment('头像') VARCHAR(300)"`
	Introduction  string    `json:"introduction" xorm:"not null default '' comment('简介') VARCHAR(300)"`
	Status        int       `json:"status" xorm:"not null default 1 comment('状态（0 停止1启动）') TINYINT(4)"`
	Utime         time.Time `json:"utime" xorm:"not null default 'CURRENT_TIMESTAMP' comment('更新时间') TIMESTAMP"`
	LastLoginTime time.Time `json:"last_login_time" xorm:"not null default '0000-00-00 00:00:00' comment('上次登录时间') DATETIME"`
	LastLoginIp   string    `json:"last_login_ip" xorm:"not null default '' comment('最近登录IP') VARCHAR(50)"`
	Ctime         time.Time `json:"ctime" xorm:"not null comment('注册时间') DATETIME"`
}

type SearchUser struct  {
	Name string `json:"name" xorm:"not null comment('姓名') VARCHAR(50)"`
}
var systemuser = "system_user"

func(u *SystemUser) GetRow() bool {
	has, err := mEngine.Get(u)
	if err==nil &&  has  {
		return true
	}
	return false
}
func (u *SystemUser) GetAll()([]SystemUser,error) {
	var systemusers []SystemUser
	err:=mEngine.Find(&systemusers)
	return systemusers,err
}
func (u *SystemUser) GetAllByName(name string)([]SearchUser,error) {
	var systemusers []SearchUser

	err:=mEngine.Table(systemuser).Where("name like ?",name+"%").Find(&systemusers)
	return systemusers,err
}

func (u *SystemUser) GetAllPage(paging *common.Paging)([]SystemUser,error) {
	var systemusers []SystemUser
	var err error
	paging.Total,err=mEngine.Where("status=?",1).Count(u)
	paging.GetPages()
	if paging.Total<1 {
		return systemusers,err
	}
	err=mEngine.Where("status=?",1).Limit(int(paging.PageSize),int(paging.StartNums)).Find(&systemusers)
	return systemusers,err
}

func (u *SystemUser) Add(roles []interface{}) (int ,error){
	session := mEngine.NewSession()
	defer session.Close()
	// add Begin() before any action
	if err := session.Begin(); err != nil {
		// if returned then will rodefer session.Close()llback automatically
		return 0,err
	}
	//var uid int64
	_,err:=session.Insert(u)
	if err!=nil {
		return 0,err
	}
	//如果没有设置权限
	if len(roles)<1 {
		return u.Id,session.Commit()
	}
	for _,k:=range roles{
		roleModel:=SystemRole{Name:k.(string)}
		has:=roleModel.GetRow()
		if !has {
			continue
		}
		if	roleModel.Status==0{
			continue
		}
		userroleModel:=SystemUserRole{SystemRoleId:roleModel.Id,SystemUserId:u.Id}
		has,err:=session.Get(&userroleModel)
		if err!=nil {
			return 0,err
		}
		if has {
			continue
		}
		_,err=session.Insert(&userroleModel)
		if err!=nil {
			return 0,err
		}
	}
	return u.Id,session.Commit()
}

func (u *SystemUser) Update(roles []interface{}) error {
	session := mEngine.NewSession()
	defer session.Close()
	// add Begin() before any action
	if err := session.Begin(); err != nil {
		// if returned then will rodefer session.Close()llback automatically
		return err
	}
	if _, err := mEngine.Where("id = ?", u.Id).Update(u); err != nil {
		return err
	}
	roleModel:=SystemUserRole{}
	if _, err := mEngine.Where("system_user_id=?",u.Id).Delete(&roleModel); err != nil {
		return err
	}
	//如果没有设置权限
	if len(roles)<1 {
		return session.Commit()
	}
	for _,k:=range roles{
		roleModel:=SystemRole{Name:k.(string)}
		has:=roleModel.GetRow()
		if !has {
			continue
		}
		if	roleModel.Status==0{
			continue
		}
		userroleModel:=SystemUserRole{SystemRoleId:roleModel.Id,SystemUserId:u.Id}
		has,err:=session.Get(&userroleModel)
		if err!=nil {
			return err
		}
		if has {
			continue
		}
		_,err=session.Insert(&userroleModel)
		if err!=nil {
			return err
		}
	}
	return session.Commit()
}
func (u *SystemUser) UpdatePasswd() error {
	if _, err := mEngine.Where("id = ?", u.Id).Update(u); err != nil {
		return err
	}
	return nil
}

func (u *SystemUser) Delete() error {
	if _, err := mEngine.Exec("update "+systemuser+" set status=? where id=?",0,u.Id); err != nil {
		return err
	}
	return nil
}
