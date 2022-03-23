package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path"
	"strings"
	"time"
	"go-admin/lib/common"
	"go-admin/lib/response"
)

func Index(c *gin.Context) {

	//menu:=systemnewdb.SystemMenu{}
	//menuArr ,_:=menu.GetRouteByUid(15)
	//menuMap :=make(map[int]string,0)
	//for _,v:=range menuArr{
	//	menuMap[v.ID]=v.Url
	//}
	//
	//menukey:=conf.Cfg.RedisPre+"menu.15"
	//rc:=cache.RedisClient.Get()
	////// 用完后将连接放回连接池
	//defer rc.Close()
	//_, err:= rc.Do("DEL", menukey)
	//
	//if err != nil {
	//	fmt.Println("redis delelte failed:", err)
	//}
	//jsonStr,err:=json.Marshal(menuMap)
	//fmt.Println(menukey)
	//_, err=rc.Do("SET",menukey,jsonStr)
	////
	//if err != nil {
	//	fmt.Println("redis set failed:", err)
	//}
	////r,err:=rc.Do("GET",menukey)
	//
	//ma ,err:=redis.String(rc.Do("GET","aaa"))
	////ma ,err:=redis.String(rc.Do("GET",menukey))
	////
	////
	//fmt.Println(ma,"aaa")
	//fmt.Println(err)
	//user := systemnewdb.SystemUser{Id:2}
	//has:=user.GetRowById()
	//fmt.Println(has)
	//response.ShowData(c, menuArr)
	//return
	c.String(http.StatusOK, "hello world")
	return
}

//上传图片
func ImgUpload(c *gin.Context) {
	//获取表单数据 参数为name值
	f, err := c.FormFile("file")
	//错误处理
	if err != nil {
		response.ShowError(c, "fail")
		return
	}
	//根据当前时间生成目录
	timestamp := time.Now().Unix()
	tm2 := time.Unix(timestamp, 0)

	//创建上传目录
	uploadDir := "upload/"
	relative := tm2.Format("20060102") + "/"
	os.MkdirAll(uploadDir+relative, os.ModePerm)

	//扩展名
	ext := path.Ext(f.Filename)
	relative = relative + common.GetRandomBoth(32) + ext

	//将文件保存至本项目根目录中
	c.SaveUploadedFile(f, uploadDir+relative)
	//保存成功返回正确的Json数据

	data := common.Url() + "/showimage?imgname=upload/" + relative
	response.ShowData(c, data)
	return
}

//删除图片
func DelImage(c *gin.Context) {
	url, has := c.GetQuery("url")
	if !has {
		response.ShowErrorParams(c, "url")
		return
	}
	url = common.SubstrContains(url, "upload/")
	//过滤危险字符
	url = strings.Replace(url, "../", "", -1)
	if common.IsFile(url) {
		err := os.Remove(url)
		if err != nil {
			response.ShowError(c, "fail")
			return
		}
	}
	response.ShowData(c, "success")
	return
}
