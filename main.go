package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"go-admin/conf"
	"go-admin/ctrl"
	"go-admin/ctrl/user"
	"go-admin/modules/cache"
	"net/http"
	"strings"
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

// 处理跨域请求,支持options访问
//func Cors() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		method := c.Request.Method
//
//		c.Header("Access-Control-Allow-Origin", "*")
//		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
//		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
//		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
//		c.Header("Access-Control-Allow-Credentials", "true")
//
//		//放行所有OPTIONS方法
//		if method == "OPTIONS" {
//			c.AbortWithStatus(http.StatusNoContent)
//		}
//		// 处理请求
//		c.Next()
//	}
//}
func Cors() gin.HandlerFunc {

	return func(c *gin.Context) {
		method := c.Request.Method

		origin := c.Request.Header.Get("Origin")
		var headerKeys []string
		for k, _ := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		fmt.Println(headerKeys)
		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}

		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Origin", "*")
			//c.Header("Access-Control-Allow-Headers", headerStr)
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
			// c.Header("Access-Control-Max-Age", "172800")
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Set("content-type", "application/json")
		}

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
	//
	//
	//
	//
	//	//return func(c *gin.Context) {
	//	//	method := c.Request.Method               //请求方法
	//	//	origin := c.Request.Header.Get("Origin") //请求头部
	//	//
	//	//	var headerKeys []string                             // 声明请求头keys
	//	//	for k, _ := range c.Request.Header {
	//	//		headerKeys = append(headerKeys, k)
	//	//	}
	//	//	headerStr := strings.Join(headerKeys, ", ")
	//	//	if headerStr != "" {
	//	//		headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
	//	//	} else {
	//	//		headerStr = "access-control-allow-origin, access-control-allow-headers"
	//	//	}
	//	//	if origin != "" {
	//	//		c.Header("Access-Control-Allow-Origin", "*")                                       // 这是允许访问所有域
	//	//		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE") //服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
	//	//		//  header的类型
	//	//		c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
	//	//		//允许跨域设置 可以返回其他子段
	//	//		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar") // 跨域关键设置 让浏览器可以解析
	//	//		c.Header("Access-Control-Max-Age", "86400")                                                                                                                                                            // 表示隔1天才发起预检请求 缓存请求信息 单位为秒
	//	//		c.Header("Access-Control-Allow-Credentials", "false")                                                                                                                                                  //	跨域请求是否需要带cookie信息 默认设置为true
	//	//		c.Set("content-type", "application/json")                                                                                                                                                              // 设置返回格式是json
	//	//	}
	//	//	//放行所有OPTIONS方法
	//	//	if method == "OPTIONS" {
	//	//		c.JSON(http.StatusOK, "Options Request!")
	//	//		//c.AbortWithStatus(http.StatusNoContent)
	//	//	}
	//	//	c.Next() //	处理请求
	//	}
	//
}
