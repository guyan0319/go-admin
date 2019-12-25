package user

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go-admin/conf"
	"go-admin/models"
	"go-admin/modules/response"
	"go-admin/public/common"
)


func Reg(c *gin.Context){
	nickname :=c.PostForm("nickname")
	passwd :=c.PostForm("passwd")
	fmt.Println(nickname)
	if nickname=="" || passwd=="" {
		response.ShowError(c,"fail")
		return
	}
	salt :=common.GetRandomBoth(4)
	passwd = common.Sha1En(passwd+salt)

}
type Userinfo struct {
	Roles []string `json:"roles"`
	Introduction string `json:"introduction"`
	Avatar string `json:"avatar"`
	Name string `json:"name"`
}
func Info(c *gin.Context){
	session := sessions.Default(c)
	v := session.Get(conf.Cfg.Token)
	if v==nil {
		response.ShowError(c,"fail")
		return
	}
	uid:=session.Get(v)
	user := models.SystemUser{Id:uid.(int)}
	has:=user.GetRow()
	if !has {
		response.ShowError(c,"user_error")
		return
	}
	userrole :=models.SystemUserRole{SystemUserId:uid.(int)}
	role,_ :=userrole.GetRowByUid()
	var info Userinfo
	info.Roles = role
	info.Name=user.Nickname
	info.Avatar=user.Avatar
	info.Introduction=user.Introduction
	response.ShowData(c, info)
	return
}
func Index(c *gin.Context)  {
	userModel:=models.SystemUser{}
	userArr, err := userModel.GetAll()
	if err != nil {
		response.ShowError(c, "fail")
		return
	}
	response.ShowData(c, userArr)
	return
}