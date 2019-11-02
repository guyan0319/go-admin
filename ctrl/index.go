package ctrl

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context)  {

	//user := models.SystemUser{Id:2}
	//has:=user.GetRowById()
	//fmt.Println(has)

	c.String(http.StatusOK, "hello world")
	return
}