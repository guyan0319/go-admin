package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"go-admin/conf"
	"go-admin/ctrl"
	"go-admin/ctrl/article"
	"go-admin/ctrl/menu"
	"go-admin/ctrl/role"
	"go-admin/ctrl/user"
	_ "go-admin/docs"
	"go-admin/models"
	"go-admin/modules/cache"
	"go-admin/modules/response"
	"go-admin/public/common"
	"net/url"
)

/*
 * 首页
 * author Guo Zhiqiang
 * datetime 2019/10/21 11:53
 */
func main() {
	Load() //载入配置
	gin.SetMode(gin.DebugMode)//开发环境
	//gin.SetMode(gin.ReleaseMode) //线上环境
	r := gin.Default()
	r.Use(cors.New(GetCorsConfig()))//跨域
	store, _ := redis.NewStoreWithPool(cache.RedisClient, []byte("secret"))
	r.Use(sessions.Sessions("gosession", store))

	//r.Use(cors.Default())//默认跨域
	r.Use(Auth())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/", ctrl.Index)
	r.POST("/upload/image", ctrl.ImgUpload)
	r.GET("/del/image", ctrl.DelImage)
	r.GET("/info", user.Info)
	r.GET("/routes",menu.List)
	r.GET("/dashboard",menu.Dashboard)
	r.GET("/role/list",menu.Roles)
	r.GET("/menu",menu.Index)
	r.POST("/menu",menu.Create)
	r.PUT("/menu",menu.Edit)
	r.DELETE("/menu",menu.Delete)
	r.GET("/user",user.Index)
	r.GET("/user/detail",user.Detail)
	r.GET("/user/search",user.Search)
	r.POST("/user/create",user.Create)
	r.POST("/user/edit",user.Edit)
	r.POST("/user/repasswd",user.Repasswd)
	r.GET("/user/delete",user.Delete)
	r.POST("/role/delete/:name",role.DeleteRole)
	r.POST("/role/update",role.UpdateRole)
	r.POST("/role/add",role.AddRole)
	r.GET("/role/index",role.Index)
	r.POST("/logout", user.Logout)
	r.POST("/login", user.Login)
	r.POST("/reg", user.Reg)
	r.POST("/articles/create", article.Create)
	r.POST("/articles/edit", article.Edit)
	r.GET("/articles/list", article.Index)
	r.GET("/articles/detail", article.Detail)
	r.GET("/showimage", article.ShowImage)

	r.GET("/pong", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8090") // listen and serve on 0.0.0.0:8080
}
func Load() {
	c := conf.Config{}
	c.Routes=[]string{"/pong","/login","/role/index","/info","/dashboard","/logout"}
	conf.Set(c)
}
func GetCorsConfig() cors.Config {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"https://admin.duiniya.com","https://admin.gzqiang.cn","http://localhost:9529","http://localhost:9528","http://localhost:9527","http://localhost"}
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
			c.Abort()
			response.ShowError(c,"nologin")
			return
		}
		uid:=session.Get(v)
		users := models.SystemUser{Id:uid.(int),Status:1}
		has:=users.GetRow()
		if !has {
			c.Abort()
			response.ShowError(c,"user_error")
			return
		}
		//特殊账号
		if users.Name==conf.Cfg.Super {
			return
		}
		menuModel:=models.SystemMenu{}
		menuMap,err:=menuModel.GetRouteByUid(uid)
		if err!=nil {
			c.Abort()
			response.ShowError(c,"unauthorized")
			return
		}
		if _,ok:=menuMap[u.Path] ;!ok{
			c.Abort()
			response.ShowError(c,"unauthorized")
			return
		}
		// access the status we are sending
		//status := c.Writer.Status()
		c.Next()
		//log.Println(status) //状态 200
	}
}
var count = 0
func test()  {
	count++
	fmt.Println(count)
}