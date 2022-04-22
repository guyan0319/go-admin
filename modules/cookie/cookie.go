package cookie

import (
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"go-admin/modules/cache"
	"go-admin/public/common"

	"strconv"
	"time"
)

var CookieName string = "GATOKEN"

var RedisExpire = 86400 * 7
var RedisExpireRepeat = 1200
var RedisRepeatSuf = "ga_redis_repeat_"

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
	return uid, err
}

//设置cookie
func SetCacheCookie(c *gin.Context, id int) (err error) {
	// 从池里获取连接
	rc := cache.RedisClient.Get()
	// 用完后将连接放回连接池
	defer rc.Close()
	cur := time.Now()
	//纳秒
	timestamps := cur.UnixNano()
	//timestamp := time.Now().Unix()
	times := strconv.FormatInt(timestamps, 10)
	key := common.Md5En(common.GetRandomString(16) + times)
	_, err = rc.Do("Set", key, id, "EX", RedisExpire)
	if err != nil {
		return
	}
	c.SetCookie(CookieName, key, RedisExpire, "/", "", false, true)
	return
}
