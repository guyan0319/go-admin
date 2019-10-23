package ctrl

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context)  {
	c.String(http.StatusOK, "hello world")
}