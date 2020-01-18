package ctrl

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"go-admin/conf"
	"go-admin/models"
	"go-admin/modules/cache"
	"go-admin/modules/response"
)

func Index(c *gin.Context)  {

	menu:=models.SystemMenu{}
	menuArr ,_:=menu.GetRowByUid(15)
	menuMap :=make(map[int]string,0)
	for _,v:=range menuArr{
		menuMap[v.Id]=v.Url
	}

	menukey:=conf.Cfg.RedisPre+"menu.15"
	rc:=cache.RedisClient.Get()
	//// 用完后将连接放回连接池
	defer rc.Close()
	_, err:= rc.Do("DEL", menukey)

	if err != nil {
		fmt.Println("redis delelte failed:", err)
	}
	jsonStr,err:=json.Marshal(menuMap)
	fmt.Println(menukey)
	_, err=rc.Do("SET",menukey,jsonStr)
	//
	if err != nil {
		fmt.Println("redis set failed:", err)
	}
	//r,err:=rc.Do("GET",menukey)

	ma ,err:=redis.String(rc.Do("GET",menukey))
	//
	//
	fmt.Println(ma,err)
	//user := models.SystemUser{Id:2}
	//has:=user.GetRowById()
	//fmt.Println(has)
	response.ShowData(c, menuArr)
	return
	//c.String(http.StatusOK, "hello world")
	//return
}