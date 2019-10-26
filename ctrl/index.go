package ctrl

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context)  {
	panic("ok")
	c.String(http.StatusOK, "hello world")
}