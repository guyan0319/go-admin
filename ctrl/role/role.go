package role

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func UpdateRole(c *gin.Context)  {
	data, _ := ioutil.ReadAll(c.Request.Body)
	fmt.Printf("ctx.Request.body: %v", string(data))




}