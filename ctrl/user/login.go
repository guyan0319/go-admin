package user

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go-admin/models"
	"go-admin/modules/response"
	"go-admin/public/common"
	"strconv"
	"time"
)
func Login(c *gin.Context) {
	nickname := c.PostForm("nickname")
	passwd := c.PostForm("passwd")
	if nickname == "" || passwd == "" {
		response.ShowError(c, "fail")
		return
	}
	user := models.SystemUser{Nickname: nickname}
	has := user.GetRowByNickname()
	if !has {
		response.ShowError(c, "fail")
		return
	}
	if common.Sha1En(passwd+user.Salt) != user.Password {
		response.ShowError(c, "fail")
		return
	}
	session := sessions.Default(c)
	var data = make(map[string]interface{}, 0)
	v := session.Get("token")
	if v == nil {
		cur := time.Now()
		//纳秒
		timestamps := cur.UnixNano()
		times := strconv.FormatInt(timestamps, 10)
		v = common.Md5En(common.GetRandomString(16) + times)
		session.Set("token", v)
		session.Save()
	}
	data["token"] = v
	response.ShowData(c, data)
	return
}
