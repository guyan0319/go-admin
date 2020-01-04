package response

import (
	"github.com/gin-gonic/gin"
	"go-admin/modules/lang"
	"net/http"
)

func ShowError(c *gin.Context, msg string){
	msg = lang.Get(msg)
	c.JSON(http.StatusOK, gin.H{
		"code": 400,
		"msg":  msg,
	})
}
func ShowErrorParams(c *gin.Context, msg string){
	msg = msg+lang.Get("not_exists")
	c.JSON(http.StatusOK, gin.H{
		"code": 400,
		"msg":  msg,
	})
}

func ShowSuccess(c *gin.Context, msg string){
	msg = lang.Get(msg)
	c.JSON(http.StatusOK, gin.H{
		"code": 20000,
		"msg":  msg,
	})
}
func ShowData(c *gin.Context, data interface{}){
	c.JSON(http.StatusOK, gin.H{
		"code": 20000,
		"data":  data,
	})
}

