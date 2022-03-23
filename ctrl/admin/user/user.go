package user

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
	"web-demo/lib/common"
	"web-demo/lib/request"
	"web-demo/lib/response"
	"web-demo/models/systemdb"
)

func Reg(c *gin.Context) {
	name := c.PostForm("name")
	passwd := c.PostForm("passwd")
	if name == "" || passwd == "" {
		response.ShowError(c, "fail")
		return
	}
	salt := common.GetRandomBoth(4)
	passwd = common.Sha1En(passwd + salt)

}

type Userinfo struct {
	Roles        []string `json:"roles"`
	Introduction string   `json:"introduction"`
	Avatar       string   `json:"avatar"`
	Name         string   `json:"name"`
}
type UserDetail struct {
	systemdb.SystemUser
	CheckedRoles []string `json:"checkedRoles"`
}

func Info(c *gin.Context) {
	session := sessions.Default(c)
	v := session.Get(common.Conf.Token)
	if v == nil {
		response.ShowError(c, "fail")
		return
	}
	uid := session.Get(v)
	db := systemdb.GetDb()

	user := systemdb.SystemUser{ID: uid.(int)}
	db.First(&user)
	userrole := systemdb.SystemUserRole{SystemUserID: uid.(int)}
	role := userrole.GetRowByUid()
	var info Userinfo
	info.Roles = role
	info.Name = user.Name
	info.Avatar = user.Avatar
	info.Introduction = user.Introduction
	response.ShowData(c, info)
	return

}

func Search(c *gin.Context) {
	name, has := c.GetQuery("name")
	if !has {
		response.ShowErrorParams(c, "name")
		return
	}
	user := systemdb.SystemUser{Name: name}
	res, _ := user.GetAllByName()
	nameList := make(map[string][]systemdb.SearchUser, 0)
	nameList["items"] = res
	response.ShowData(c, nameList)
	return
}
func Detail(c *gin.Context) {
	id, has := c.GetQuery("id")
	if !has {
		response.ShowErrorParams(c, "id")
		return
	}
	user := systemdb.SystemUser{}
	user.ID, _ = strconv.Atoi(id)
	res, _ := user.GetRow()
	if res < 1 {
		response.ShowError(c, "user_error")
		return
	}
	userrole := systemdb.SystemUserRole{SystemUserID: user.ID}
	role := userrole.GetRowByUid()

	detail := UserDetail{}
	detail.CheckedRoles = role
	detail.ID = user.ID
	detail.Name = user.Name
	detail.Nickname = user.Nickname
	detail.Phone = user.Phone
	detail.State = user.State
	response.ShowData(c, detail)
	return
}
func Delete(c *gin.Context) {
	id, has := c.GetQuery("id")
	if !has {
		response.ShowErrorParams(c, "id")
		return
	}
	user := systemdb.SystemUser{}
	user.ID, _ = strconv.Atoi(id)
	err := user.Delete()
	if err != nil {
		response.ShowError(c, "fail")
		return
	}
	response.ShowData(c, "success")
	return
}
func Index(c *gin.Context) {
	page, _ := strconv.ParseInt(c.Query("page"), 10, 64)
	limit, _ := strconv.ParseInt(c.Query("limit"), 10, 64)

	paging := &common.Paging{Page: page, PageSize: limit}
	userModel := systemdb.SystemUser{}
	userArr, err := userModel.GetAllPage(paging)
	if err != nil {
		response.ShowError(c, "fail")
		return
	}
	data := make(map[string]interface{})
	data["items"] = userArr
	data["total"] = paging.Total
	response.ShowData(c, data)
	return
}
func Create(c *gin.Context) {
	data, err := request.GetJson(c)
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
	if _, ok := data["state"]; !ok {
		response.ShowError(c, "fail")
		return
	}
	userModel := systemdb.SystemUser{}
	userModel.Name = data["name"].(string)
	has, _ := userModel.GetRow()
	if has > 0 {
		response.ShowError(c, "name_exists")
		return
	}
	userModel.Password = data["password"].(string)
	if userModel.Password != data["repassword"].(string) {
		response.ShowError(c, "fail")
		return
	}
	userModel.Salt = common.GetRandomBoth(4)
	userModel.Password = common.Sha1En(userModel.Password + userModel.Salt)
	userModel.Name = data["name"].(string)
	userModel.Nickname = data["nickname"].(string)
	if _, ok := data["phone"]; ok {
		userModel.Phone = data["phone"].(string)
	}
	if _, ok := data["state"]; ok && data["state"].(bool) {
		userModel.State = 1
	}
	userModel.Avatar = "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif"
	userModel.Ctime = time.Now()
	if _, ok := data["checkedRoles"]; !ok {
		response.ShowError(c, "fail")
		return
	}
	roles := data["checkedRoles"].([]interface{})

	err = userModel.Add(roles)
	if err != nil {
		response.ShowError(c, "fail")
		return
	}
	response.ShowData(c, userModel)
	return
}
func Edit(c *gin.Context) {
	data, err := request.GetJson(c)
	if err != nil {
		response.ShowError(c, "fail")
		return
	}
	if _, ok := data["id"]; !ok {
		response.ShowError(c, "fail")
		return
	}
	userModel := systemdb.SystemUser{}
	userModel.ID = int(data["id"].(float64))
	has, _ := userModel.GetRow()
	if has < 1 {
		response.ShowError(c, "user_error")
		return
	}
	if _, ok := data["nickname"]; !ok {
		response.ShowError(c, "fail")
		return
	}
	if _, ok := data["state"]; !ok {
		response.ShowError(c, "fail")
		return
	}
	if _, ok := data["state"]; ok && data["state"].(bool) {
		userModel.State = 1
	} else {
		userModel.State = 0
	}
	userModel.Nickname = data["nickname"].(string)
	if _, ok := data["phone"]; ok {
		userModel.Phone = data["phone"].(string)
	}
	if _, ok := data["checkedRoles"]; !ok {
		response.ShowError(c, "fail")
		return
	}
	roles := data["checkedRoles"].([]interface{})
	err = userModel.Update(roles)
	if err != nil {
		response.ShowError(c, "fail")
		return
	}
	response.ShowData(c, userModel)
	return
}
func Repasswd(c *gin.Context) {
	data, err := request.GetJson(c)
	if err != nil {
		response.ShowError(c, "fail")
		return
	}
	if _, ok := data["id"]; !ok {
		response.ShowError(c, "fail")
		return
	}
	userModel := systemdb.SystemUser{}
	userModel.ID = int(data["id"].(float64))
	has, _ := userModel.GetRow()
	if has < 1 {
		response.ShowError(c, "user_error")
		return
	}
	if userModel.Name == "admin" {
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
	userModel.Password = data["password"].(string)
	if userModel.Password != data["repassword"].(string) {
		response.ShowError(c, "fail")
		return
	}
	userModel.Salt = common.GetRandomBoth(4)
	userModel.Password = common.Sha1En(userModel.Password + userModel.Salt)
	err = userModel.UpdatePasswd()
	if err != nil {
		response.ShowError(c, "fail")
		return
	}
	response.ShowData(c, userModel)
	return
}
