package user

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go-admin/conf"
	"go-admin/models"
	"go-admin/modules/request"
	"go-admin/modules/response"
	"go-admin/public/common"
	"strconv"
	"time"
)

func Reg(c *gin.Context){
	name :=c.PostForm("name")
	passwd :=c.PostForm("passwd")
	if name=="" || passwd=="" {
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
type UserDetail struct {
	models.SystemUser
	CheckedRoles []string `json:"checkedRoles"`
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
	info.Name=user.Name
	info.Avatar=user.Avatar
	info.Introduction=user.Introduction
	response.ShowData(c, info)
	return
}
func Search(c *gin.Context)  {
	name,has:=c.GetQuery("name")
	if	!has{
		response.ShowErrorParams(c, "name")
		return
	}
	user := models.SystemUser{}
	res ,_:=user.GetAllByName(name)
	nameList :=make(map[string][]models.SearchUser,0)
	nameList["items"]=res
	response.ShowData(c, nameList)
	return
}
func Detail(c *gin.Context){
	id,has:=c.GetQuery("id")
	if	!has{
		response.ShowErrorParams(c, "id")
		return
	}
	user := models.SystemUser{}
	user.Id,_=strconv.Atoi(id)
	has=user.GetRow()
	if !has {
		response.ShowError(c,"user_error")
		return
	}
	userrole :=models.SystemUserRole{SystemUserId:user.Id}
	role,_ :=userrole.GetRowByUid()

	detail:=UserDetail{}
	detail.CheckedRoles=role
	detail.Id=user.Id
	detail.Name=user.Name
	detail.Nickname=user.Nickname
	detail.Phone=user.Phone
	detail.Status=user.Status
	response.ShowData(c, detail)
	return
}
func Delete(c *gin.Context){
	id,has:=c.GetQuery("id")
	if	!has{
		response.ShowErrorParams(c, "id")
		return
	}
	user := models.SystemUser{}
	user.Id,_=strconv.Atoi(id)
	err:=user.Delete()
	if err!=nil {
		response.ShowError(c,"fail")
		return
	}
	response.ShowData(c,"success")
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
	userModel := models.SystemUser{};
	userModel.Name=data["name"].(string)
	has:=userModel.GetRow()
	if has {
		response.ShowError(c, "name_exists")
		return
	}
	userModel.Password=data["password"].(string)
	if	userModel.Password!=data["repassword"].(string){
		response.ShowError(c, "fail")
		return
	}
	userModel.Salt =common.GetRandomBoth(4)
	userModel.Password = common.Sha1En(userModel.Password+userModel.Salt)
	userModel.Name = data["name"].(string)
	userModel.Nickname = data["nickname"].(string)
	if _, ok := data["phone"]; ok {
		userModel.Phone = data["phone"].(string)
	}
	if _, ok := data["status"]; ok && data["status"].(bool) {
		userModel.Status=1
	}
	userModel.Avatar="https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif"
	userModel.Ctime=time.Now()
	if _, ok := data["checkedRoles"]; !ok {
		response.ShowError(c, "fail")
		return
	}
	roles:=data["checkedRoles"].([]interface{})
	_,err=userModel.Add(roles)
	if err!=nil {
		response.ShowError(c,"fail")
		return
	}
	response.ShowData(c,userModel)
	return
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
	userModel:=models.SystemUser{}
	userModel.Id=int(data["id"].(float64))
	has:=userModel.GetRow()
	if !has {
		response.ShowError(c,"user_error")
		return
	}
	if _, ok := data["nickname"]; !ok {
		response.ShowError(c, "fail")
		return
	}
	if _, ok := data["status"]; !ok {
		response.ShowError(c, "fail")
		return
	}
	if _, ok := data["status"]; ok && data["status"].(bool) {
		userModel.Status=1
	}else{
		userModel.Status=0
	}
	userModel.Nickname = data["nickname"].(string)
	if _, ok := data["phone"]; ok {
		userModel.Phone = data["phone"].(string)
	}
	if _, ok := data["checkedRoles"]; !ok {
		response.ShowError(c, "fail")
		return
	}
	roles:=data["checkedRoles"].([]interface{})
	err=userModel.Update(roles)
	if err!=nil {
		response.ShowError(c,"fail")
		return
	}
	response.ShowData(c,userModel)
	return
}
func Repasswd(c *gin.Context)  {
	data,err:=request.GetJson(c)
	if err != nil {
		response.ShowError(c, "fail")
		return
	}
	if _, ok := data["id"]; !ok {
		response.ShowError(c, "fail")
		return
	}
	userModel:=models.SystemUser{}
	userModel.Id=int(data["id"].(float64))
	has:=userModel.GetRow()
	if !has {
		response.ShowError(c,"user_error")
		return
	}
	if userModel.Name=="admin" {
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
	userModel.Password=data["password"].(string)
	if	userModel.Password!=data["repassword"].(string){
		response.ShowError(c, "fail")
		return
	}
	userModel.Salt =common.GetRandomBoth(4)
	userModel.Password = common.Sha1En(userModel.Password+userModel.Salt)
	err=userModel.UpdatePasswd()
	if err!=nil {
		response.ShowError(c,"fail")
		return
	}
	response.ShowData(c,userModel)
	return
}
