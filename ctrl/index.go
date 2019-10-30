package ctrl

import (
	"github.com/gin-gonic/gin"
	"go-admin/modules/response"
	"net/http"
)

func Index(c *gin.Context)  {
	//uid,err :=cookie.GetCacheCookie(c)
	//if err!=nil{
	//	fmt.Println(err)
	//	response.ShowError(c,"fail")
	//	return
	//}
	uid :=c.Query("username")
	response.ShowData(c,uid)
	c.String(http.StatusOK, "hello world")
	return
}