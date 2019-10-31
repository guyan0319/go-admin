package user

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go-admin/conf"
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
	fmt.Println(salt)
	fmt.Println(passwd)
}
func Info(c *gin.Context){
	session := sessions.Default(c)
	v := session.Get(conf.Cfg.Token)
	fmt.Println(v)
	data:="ok"
	response.ShowData(c, data)
	return
}