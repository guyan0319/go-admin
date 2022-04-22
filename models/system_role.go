package models

import (
	"fmt"
	"strings"
	"time"
)

type SystemRole struct {
	Id          int       `json:"id" xorm:"not null pk autoincr comment('主键') INT(11)"`
	Name        string    `json:"name" xorm:"not null comment('角色名称') VARCHAR(100)"`
	AliasName   string    `json:"alias_name" xorm:"not null default '' comment('别名') VARCHAR(50)"`
	Description string    `json:"description" xorm:"not null default '' comment('描述') VARCHAR(200)"`
	Status      int       `json:"status" xorm:"not null default 1 comment('角色状态（0无效1有效）') index TINYINT(4)"`
	Type        int       `json:"type" xorm:"not null default 1 comment('属于哪个应用') index INT(4)"`
	Ctime       time.Time `json:"ctime" xorm:"not null comment('创建时间') DATETIME"`
}

var systemrole = "system_role"

func (r *SystemRole) GetRow() bool {
	has, err := mEngine.Get(r)
	if err == nil && has {
		return true
	}
	return false
}
func (r *SystemRole) GetRowByName() bool {
	role := SystemRole{}
	has, err := mEngine.Where("name = ?", r.Name).Get(&role)
	if err == nil && has {
		return true
	}
	return false
}

func (r *SystemRole) Update(data []int) error {
	session := mEngine.NewSession()

	// add Begin() before any action
	if err := session.Begin(); err != nil {
		return err
	}
	if _, err := session.Where("id = ?", r.Id).Update(r); err != nil {
		fmt.Println(err)
		return err
	}
	if len(data)<=0 {
		return  session.Commit()
	}
	rolemenu:=SystemRoleMenu{SystemRoleId:r.Id}
	if _, err := session.Delete(&rolemenu); err != nil {
		fmt.Println(err)
		return err
	}
	for _,value:=range data {
		rm:=SystemRoleMenu{SystemRoleId:r.Id,SystemMenuId:value}
		if _, err := session.Insert(&rm); err != nil {
			fmt.Println(err)
			return err
		}
	}
	// add Commit() after all actions
	return  session.Commit()
}
func (r *SystemRole) Updateold(data []interface{}) error {
	session := mEngine.NewSession()
	defer session.Close()
	// add Begin() before any action
	if err := session.Begin(); err != nil {
		// if returned then will rollback automatically
		return err
	}
	if _, err := session.Where("id = ?", r.Id).Update(r); err != nil {
		fmt.Println(err)
		return err
	}
	if len(data)<=0 {
		return  session.Commit()
	}
	rolemenu:=SystemRoleMenu{SystemRoleId:r.Id}
	if _, err := session.Delete(&rolemenu); err != nil {
		fmt.Println(err)
		return err
	}
	for _,value:=range data  {
		menu:=SystemMenu{}

		pathMain:=value.(map[string]interface{})["path"].(string)
		menu.Path=pathMain
		menu.Component=value.(map[string]interface{})["component"].(string)
		menu.Type=2
		initMenu:=SystemMenu{}
		has:=initMenu.GetRowByPathCT(menu)
		if !has {
			continue
		}
		rm:=SystemRoleMenu{SystemRoleId:r.Id,SystemMenuId:initMenu.Id}
		if _, err := session.Insert(&rm); err != nil {
			fmt.Println(err)
			return err
		}
		children :=value.(map[string]interface{})["children"]
		if children==nil {
			continue
		}
		for _,v:=range children.([]interface{})  {
			menu:=SystemMenu{}
			menu.Path=strings.TrimPrefix(v.(map[string]interface{})["path"].(string),pathMain+"/")
			menu.Component=v.(map[string]interface{})["component"].(string)
			menu.Type=2
			initMenu:=SystemMenu{}
			has:=initMenu.GetRowByPathCT(menu)
			if !has {
				continue
			}
			rm:=SystemRoleMenu{SystemRoleId:r.Id,SystemMenuId:initMenu.Id}
			if _, err := session.Insert(&rm); err != nil {
				fmt.Println(err)
				return err
			}
		}
	}
	// add Commit() after all actions
	return  session.Commit()
}
func (r *SystemRole) AddCommit(data []interface{}) error {
	session := mEngine.NewSession()
	defer session.Close()
	// add Begin() before any action
	if err := session.Begin(); err != nil {
		// if returned then will rollback automatically
		return err
	}
	if _, err := session.Insert(r); err != nil {
		fmt.Println(err)
		return err
	}
	if len(data)<=0 {
		return  session.Commit()
	}
	for _,value:=range data  {
		menu:=SystemMenu{}
		pathMain:=value.(map[string]interface{})["path"].(string)
		menu.Path=pathMain
		menu.Component=value.(map[string]interface{})["component"].(string)
		menu.Type=2
		initMenu:=SystemMenu{}
		has:=initMenu.GetRowByPathCT(menu)
		if !has {
			continue
		}
		rm:=SystemRoleMenu{SystemRoleId:r.Id,SystemMenuId:initMenu.Id}
		if _, err := session.Insert(&rm); err != nil {
			fmt.Println(err)
			return err
		}
		children :=value.(map[string]interface{})["children"]
		if children==nil {
			continue
		}
		for _,v:=range children.([]interface{})  {
			menu:=SystemMenu{}
			strings.TrimPrefix(v.(map[string]interface{})["path"].(string),pathMain+"/")
			menu.Component=v.(map[string]interface{})["component"].(string)
			menu.Type=2
			initMenu:=SystemMenu{}
			has:=initMenu.GetRowByPathCT(menu)
			if !has {
				continue
			}
			rm:=SystemRoleMenu{SystemRoleId:r.Id,SystemMenuId:initMenu.Id}
			if _, err := session.Insert(&rm); err != nil {
				fmt.Println(err)
				return err
			}
		}
	}
	// add Commit() after all actions
	return  session.Commit()
}
func (r *SystemRole) Add() bool {
	if r.Name == "" {
		return false
	}
	role := SystemRole{}
	has, err := mEngine.Where("name = ?", r.Name).Desc("id").Get(&role)
	if err != nil {
		fmt.Println(err)
		return false
	}
	if has {
		r.Id=role.Id
		return true
	}
	r.Status=1
	r.Type=1
	_, err = mEngine.Insert(r)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
func  (r *SystemRole) GetRowMenu()(map[int][]string){
	var sr []SystemRole
	err:=mEngine.Find(&sr)
	if err!=nil {
		panic(err)
	}
	var srMap map[int]string
	srMap = make(map[int]string,0)
	for _,v:=range sr{
		srMap[v.Id]=v.Name
	}
	var srm=SystemRoleMenu{}
	rmArr,_:=srm.GetAll()
	var mrMap=make(map[int][]string,0)
	for _,value:=range rmArr{
		_,ok:=srMap[value.SystemRoleId]
		if ok {
			mrMap[value.SystemMenuId]=append(mrMap[value.SystemMenuId],srMap[value.SystemRoleId])
		}
	}
	return mrMap
}

func  (r *SystemRole) GetAll()([]SystemRole){
	var sr []SystemRole
	err:=mEngine.Find(&sr)
	if err!=nil {
		panic(err)
	}
	return sr
}
func(r *SystemRole) Delete()(error) {
	session := mEngine.NewSession()
	defer session.Close()
	// add Begin() before any action
	if err := session.Begin(); err != nil {
		// if returned then will rollback automatically
		return err
	}
	//rolemenu:=SystemRoleMenu{SystemRoleId:r.Id}
	//if _, err := session.Delete(&rolemenu); err != nil {
	//	fmt.Println(err)
	//	return err
	//}
	//roleuser:=SystemUserRole{SystemRoleId:r.Id}
	//if _, err := session.Delete(&roleuser); err != nil {
	//	fmt.Println(err)
	//	return err
	//}
	if _, err := session.Exec("update "+systemrole+" set status=? where id=?",0,r.Id); err != nil {
		return err
	}
	//if _, err := session.Delete(r); err != nil {
	//	return err
	//}
	// add Commit() after all actions
	return  session.Commit()

}
func(r *SystemRole) GetNameList()([]string) {
	var list []string
	err:=mEngine.Table(systemrole).Where("status=?",1).Cols("name").Find(&list)
	if err!=nil {
		panic(err)
	}
	return list
}