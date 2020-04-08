package article

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-admin/models"
	"go-admin/modules/request"
	"go-admin/modules/response"
	"go-admin/public/common"
	"strconv"
)

func Create(c *gin.Context)  {

	data,err:=request.GetJson(c)
	if err != nil {
		response.ShowError(c, "fail")
		return
	}
	if _, ok := data["status"]; !ok {
		response.ShowError(c, "fail")
		return
	}
	if _, ok := data["title"]; !ok {
		response.ShowError(c, "fail")
		return
	}
	if _, ok := data["content"]; !ok {
		response.ShowError(c, "fail")
		return
	}
	if _, ok := data["content_short"]; !ok {
		response.ShowError(c, "fail")
		return
	}
	if _, ok := data["source_uri"]; !ok {
		response.ShowError(c, "fail")
		return
	}
	if _, ok := data["image_uri"]; !ok {
		response.ShowError(c, "fail")
		return
	}
	if _, ok := data["display_time"]; !ok {
		response.ShowError(c, "fail")
		return
	}
	if _, ok := data["comment_disabled"]; !ok {
		response.ShowError(c, "fail")
		return
	}
	if _, ok := data["importance"]; !ok {
		response.ShowError(c, "fail")
		return
	}
	if _, ok := data["author"]; !ok {
		response.ShowError(c, "fail")
		return
	}
	if data["title"].(string) ==""{
		response.ShowErrorParams(c, "title")
		return
	}
	if data["image_uri"].(string) ==""{
		response.ShowErrorParams(c, "image_uri")
		return
	}
	//fmt.Println(data)
	model :=models.SystemArticle{}
	model.ImageUrl,_=common.WriteFile("./upload",data["image_uri"].(string))
	model.Title=data["title"].(string)
	model.Content=data["content"].(string)
	model.ContentShort=data["content_short"].(string)
	model.SourceUrl=data["source_uri"].(string)

	model.Importance=int(data["importance"].(float64))
	status:=int(data["status"].(float64))
	if status ==1{
		model.Status=1
	}
	if data["comment_disabled"].(bool) {
		model.CommentDisabled =1
	}
	model.Ptime = common.StrToTimes(data["display_time"].(string))
	model.Ctime=int(model.Ptime.Unix())
	model.Mtime=model.Ptime
	user :=models.SystemUser{Name:data["author"].(string)}
	has :=user.GetRow()
	if !has {
		common.ShowMsg("fail")
		return
	}
	model.Author=user.Id
	_,err=model.Add()
	if err!=nil {
		fmt.Println(err)
		common.ShowMsg("fail")
		return
	}
	response.ShowData(c,model)
	return
}

func Index(c *gin.Context)  {
	page, _ := strconv.ParseInt(c.Query("page"), 10, 64)
	limit, _ := strconv.ParseInt(c.Query("limit"), 10, 64)

	paging:=&common.Paging{Page:page,PageSize:limit}
	articleModel:=models.SystemArticle{}
	articleArr, err := articleModel.GetAllPage(paging)
	var articlePageArr []models.SystemArticlePage

	for _,v :=range articleArr{
		userModel := models.SystemUser{}
		userModel.Id=v.Author
		has :=userModel.GetRow()
		if !has {
			continue
		}
		fmt.Println("aa")
		articlePageArr = append(articlePageArr,models.SystemArticlePage{SystemArticle:v,AuthorName:userModel.Name})
	}
	if err != nil {
		response.ShowError(c, "fail")
		return
	}
	data :=make(map[string]interface{})
	data["items"]=articlePageArr
	data["total"]=paging.Total
	response.ShowData(c, data)
	return
}