package menu

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go-admin/conf"
	"go-admin/modules/response"
)

func List(c *gin.Context)  {
	session := sessions.Default(c)
	v := session.Get(conf.Cfg.Token)
	if v==nil {
		response.ShowError(c,"fail")
		return
	}
	fmt.Println(v)
	




}