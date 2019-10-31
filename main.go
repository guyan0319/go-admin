package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"go-admin/conf"
	"go-admin/ctrl"
	"go-admin/ctrl/user"
	"go-admin/modules/cache"
)

/*
 * 首页
 * author Guo Zhiqiang
 * datetime 2019/10/21 11:53
 */
func main() {
	Load() //载入配置
	//gin.SetMode(gin.DebugMode)//开发环境
	gin.SetMode(gin.ReleaseMode) //线上环境
	r := gin.Default()
	store, _ := redis.NewStoreWithPool(cache.RedisClient, []byte("secret"))
	r.Use(sessions.Sessions("gosession", store))
	r.Use(cors.New(GetCorsConfig()))
	//r.Use(cors.Default())//默认跨域
	r.GET("/", ctrl.Index)
	r.POST("/login", user.Login)
	r.POST("/reg", user.Reg)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8090") // listen and serve on 0.0.0.0:8080
}
func Load() {
	c := conf.Config{}
	conf.Set(c)

}
func GetCorsConfig() cors.Config {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost"}
	config.AllowMethods = []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"}
	config.AllowCredentials = true
	config.AllowHeaders = []string{"x-requested-with", "Content-Type", "AccessToken", "X-CSRF-Token", "Authorization"}
	return config
}
