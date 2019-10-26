package ctrl

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-admin/modules/cookie"
	"go-admin/modules/response"
	"net/http"
)

func Index(c *gin.Context)  {
	uid,err :=cookie.GetCacheCookie(c)
	if err!=nil{
		fmt.Println(err)
		response.ShowError(c,"fail")
		return
	}
	response.ShowData(c,uid)
	c.String(http.StatusOK, "hello world")
	return
}