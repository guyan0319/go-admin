package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
)
type User struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password   string `form:"password" json:"password" binding:"required"`
}
func Login(c *gin.Context) {
	var u User
	err :=c.BindJSON(&u)
	fmt.Println(err)
	fmt.Println(u)
	//nickname := c.PostForm("username")
	//passwd := c.PostForm("password")
	//param3,_ :=c.GetPostForm("username")
	//fmt.Println(param3)
	//if nickname == "" || passwd == "" {
	//	response.ShowError(c, "fail")
	//	return
	//}
	//user := models.SystemUser{Nickname: nickname}
	//has := user.GetRowByNickname()
	//if !has {
	//	response.ShowError(c, "fail")
	//	return
	//}
	//if common.Sha1En(passwd+user.Salt) != user.Password {
	//	response.ShowError(c, "fail")
	//	return
	//}
	////session := sessions.Default(c)
	//var data = make(map[string]interface{}, 0)
	//v := session.Get("token")
	//if v == nil {
	//	cur := time.Now()
	//	//纳秒
	//	timestamps := cur.UnixNano()
	//	times := strconv.FormatInt(timestamps, 10)
	//	v = common.Md5En(common.GetRandomString(16) + times)
	//	session.Set("token", v)
	//	_=session.Save()
	//}
	//data["token"] = v
	//response.ShowData(c, data)
	return
}
