package user

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
	"go-admin/lib/cache"
	"go-admin/lib/common"
	"go-admin/lib/response"
	"go-admin/models/systemnewdb"
)

type User struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// @登录
// @Summary 登录
// @Tags user
// @Description login
// @Accept  json
// @Produce json
// @Param   username     path    string     true        "username"
// @Param   passwd     path    string     true        "passwd"
// @Success 200 {string} string	"ok"
// @Router /login [post]
//func Login(c *gin.Context) {
//
//	session := sessions.Default(c)
//	var data = make(map[string]interface{}, 0)
//	v := session.Get(common.Conf.Token)
//
//	if v == nil {
//		cur := time.Now()
//		//纳秒
//		timestamps := cur.UnixNano()
//		times := strconv.FormatInt(timestamps, 10)
//		v = common.Md5En(common.GetRandomString(16) + times)
//		session.Set(common.Conf.Token, v)
//		session.Set(v, 1)
//		session.Save()
//		fmt.Println("设置成功")
//	}
//	data[common.Conf.Token] = v
//	response.ShowData(c, data)
//	return
//}

func Login(c *gin.Context) {
	var u User
	err := c.BindJSON(&u)
	if err != nil {
		response.ShowError(c, "fail")
		return
	}
	if u.Username == "" || u.Password == "" {
		response.ShowError(c, "fail")
		return
	}
	user := systemnewdb.SystemUser{Name: u.Username}
	has, _ := user.GetRow()
	if has < 1 {
		response.ShowError(c, "fail")
		return
	}
	if common.Sha1En(u.Password+user.Salt) != user.Password {
		response.ShowError(c, "fail")
		return
	}
	session := sessions.Default(c)
	var data = make(map[string]interface{}, 0)
	v := session.Get(common.Conf.Token)

	if v == nil {
		cur := time.Now()
		//纳秒
		timestamps := cur.UnixNano()
		times := strconv.FormatInt(timestamps, 10)
		v = common.Md5En(common.GetRandomString(16) + times)
		session.Set(common.Conf.Token, v)
		session.Set(v, user.ID)
		err = session.Save()
		fmt.Println("设置成功")
	}
	data[common.Conf.Token] = v
	response.ShowData(c, data)
	return
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	v := session.Get(common.Conf.Token)
	id := session.Get(v)
	strId := strconv.Itoa(id.(int))
	menukey := common.Conf.RedisPre + "menu." + strId
	rc := cache.RedisClient.Get()
	////// 用完后将连接放回连接池
	defer rc.Close()
	_, err := rc.Do("DEL", menukey)
	if err != nil {
		fmt.Println("redis delelte failed:", err)
	}
	session.Clear()
	//清除session
	_ = session.Save()
	response.ShowSuccess(c, "success")
	return
}
