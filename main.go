package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"go-admin/conf"
	"go-admin/ctrl"
	"go-admin/ctrl/menu"
	"go-admin/ctrl/role"
	"go-admin/ctrl/user"
	_ "go-admin/docs"
	"go-admin/models"
	"go-admin/modules/cache"
	"go-admin/modules/response"
	"go-admin/public/common"
	"log"
	"net/url"
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
	r.Use(cors.New(GetCorsConfig()))//跨域
	//r.Use(cors.Default())//默认跨域
	r.Use(Auth())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/", ctrl.Index)
	r.GET("/info", user.Info)
	r.GET("/routes",menu.List)
	r.GET("/dashboard",menu.Dashboard)
	r.GET("/role/list",menu.Roles)
	r.GET("/menu",menu.Index)
	r.POST("/menu",menu.Create)
	r.PUT("/menu",menu.Edit)
	r.DELETE("/menu",menu.Delete)
	r.GET("/user",user.Index)
	r.GET("/user/create",user.Create)
	r.POST("/user/edit",user.Edit)
	r.POST("/role/delete/:name",role.DeleteRole)
	r.POST("/role/update",role.UpdateRole)
	r.POST("/role/add",role.AddRole)
	r.POST("/logout", user.Logout)
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
	c.Routes=[]string{"/login"}
	conf.Set(c)

}
func GetCorsConfig() cors.Config {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:9529","http://localhost:9528","http://localhost:9527","http://localhost"}
	config.AllowMethods = []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"}
	config.AllowCredentials = true
	config.AllowHeaders = []string{"x-requested-with", "Content-Type", "AccessToken", "X-CSRF-Token","X-Token", "Authorization","token"}
	return config
}
func Auth() gin.HandlerFunc{
	return func(c *gin.Context) {
		u,err:= url.Parse(c.Request.RequestURI)
		if err != nil {
			panic(err)
		}
		if common.InArrayString(u.Path,&conf.Cfg.Routes) {
			c.Next()
			return
		}
		session := sessions.Default(c)
		v := session.Get(conf.Cfg.Token)
		if v==nil {
			response.ShowError(c,"nologin")
			return
		}
		uid:=session.Get(v)
		user := models.SystemUser{Id:uid.(int),Status:1}
		has:=user.GetRow()
		if !has {
			response.ShowError(c,"user_error")
			return
		}
		constant:=models.SystemMenu{Type:1}
		constant.GetRow()
		spcial:=models.SystemMenu{Type:3}
		spcial.GetRow()

		//特殊账号
		if user.Name==conf.Cfg.Super {



		}


		//t := time.Now()
		//// Set example variable
		//c.Set("example", "12345")
		// before request
		//c.Next()
		//// after request
		//latency := time.Since(t)
		//log.Print(latency) //时间  0s
		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status) //状态 200
	}
}