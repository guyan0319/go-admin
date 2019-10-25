package cookie

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-admin/modules/cache"

	"strconv"
	"time"
)

//获取cookie
func GetCacheCookie(c *gin.Context) (int64, error) {
	val, err := c.Cookie(CookieName)
	if err != nil {
		return 0, err
	}
	// 从池里获取连接

	rc := cache.RedisClient.Get()
	// 用完后将连接放回连接池
	defer rc.Close()
	uid, err := redis.Int64(rc.Do("GET", val))
	if err != nil {
		return 0, err
	}
	fmt.Println(uid)
	return uid, err
}

//设置cookie
func SetCacheCookie(c *gin.Context, id int64) (err error) {
	// 从池里获取连接
	rc := lib.RedisClient.Get()
	// 用完后将连接放回连接池
	defer rc.Close()
	cur := time.Now()
	//纳秒
	timestamps := cur.UnixNano()
	//timestamp := time.Now().Unix()
	times := strconv.FormatInt(timestamps, 10)
	key := lib.Md5En(lib.GetRandomString(16) + times)
	expire := 3600 * 24 * 7 //有效期7天
	_, err = rc.Do("Set", key, id, "EX", expire)
	if err != nil {
		return
	}
	c.SetCookie(CookieName, key, expire, "/", "", false, true)
	return
}