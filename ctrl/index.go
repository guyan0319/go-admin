package ctrl

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context)  {
	session := sessions.Default(c)
	v:=session.Get("token")
	if v==nil {
		//fmt.Println("设置成功")
		//session.Set("token","tokens")
		//err := session.Save()
		//fmt.Println(err)
	}else{
		fmt.Println("token:",v)
	}
	c.String(http.StatusOK, "hello world")
	return
}