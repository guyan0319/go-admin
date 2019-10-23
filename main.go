package main

import (
	"github.com/gin-gonic/gin"
)

/*
 *   首页
 * author Guo Zhiqiang
 * datetime 2019/10/21 11:53
 */
func main() {
	r := gin.Default()
	r.GET("/",ctrl.Index)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8090") // listen and serve on 0.0.0.0:8080

}
