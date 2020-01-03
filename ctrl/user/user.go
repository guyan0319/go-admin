package user

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go-admin/conf"
	"go-admin/models"
	"go-admin/modules/request"
	"go-admin/modules/response"
	"go-admin/public/common"
	"strconv"
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
	page, _ := strconv.ParseInt(c.Query("page"), 10, 64)
	limit, _ := strconv.ParseInt(c.Query("limit"), 10, 64)

	paging:=&common.Paging{Page:page,PageSize:limit}
	userModel:=models.SystemUser{}
	userArr, err := userModel.GetAllPage(paging)
	if err != nil {
		response.ShowError(c, "fail")
		return
	}
	data :=make(map[string]interface{})
	data["items"]=userArr
	data["total"]=paging.Total

	response.ShowData(c, data)
	return
}
func Create(c *gin.Context)  {
	data,err:=request.GetJson(c)
	if err != nil {
		response.ShowError(c, "fail")
		return
	}
	if _, ok := data["name"]; !ok {
		response.ShowError(c, "fail")
		return
	}
	if _, ok := data["nickname"]; !ok {
		response.ShowError(c, "fail")
		return
	}
	if _, ok := data["password"]; !ok {
		response.ShowError(c, "fail")
		return
	}
	if _, ok := data["repassword"]; !ok {
		response.ShowError(c, "fail")
		return
	}
	if _, ok := data["status"]; !ok {
		response.ShowError(c, "fail")
		return
	}
	if _, ok := data["avatar"]; !ok {
		response.ShowError(c, "fail")
		return
	}
	userModel := models.SystemUser{};
	userModel.Nickname=data["name"].(string)
	has:=userModel.GetRow()
	if has {
		response.ShowError(c, "name_exists")
		return
	}
	userModel.Password=data["password"].(string)
	if	userModel.Password!=data["repassword"].(string){
		response.ShowError(c, "nickname_exists")
		return
	}
	userModel.Salt =common.GetRandomBoth(4)
	userModel.Password = common.Sha1En(userModel.Password+userModel.Salt)
	userModel.Name = data["name"].(string)
	if _, ok := data["status"]; ok {

	}

}
func Edit(c *gin.Context)  {
	data,err:=request.GetJson(c)
	if err != nil {
		response.ShowError(c, "fail")
		return
	}
	if _, ok := data["id"]; !ok {
		response.ShowError(c, "fail")
		return
	}
}