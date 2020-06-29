package article

import (
	"github.com/gin-gonic/gin"
	"go-admin/conf"
	"go-admin/models"
	"go-admin/modules/request"
	"go-admin/modules/response"
	"go-admin/public/common"
	"strconv"
	"strings"
	"time"
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
	//model.ImageUri,_=common.WriteFile("./upload",data["image_uri"].(string))
	model.ImageUri=common.Substr(data["image_uri"].(string),"upload/")
	model.Title=data["title"].(string)
	model.Content=data["content"].(string)
	//model.Content=common.Base64Content(conf.Cfg.Host+"/showimage?imgname=upload/","./upload",model.Content)
	model.ContentShort=data["content_short"].(string)
	model.SourceUri=data["source_uri"].(string)

	model.Importance=int(data["importance"].(float64))
	status:=int(data["status"].(float64))
	if status ==1{
		model.Status=1
	}
	if data["comment_disabled"].(bool) {
		model.CommentDisabled =1
	}
	model.Display_time = common.StrToTimes(data["display_time"].(string))
	model.Ctime=int(model.Display_time.Unix())
	model.Mtime=model.Display_time
	user :=models.SystemUser{Name:data["author"].(string)}
	has :=user.GetRow()
	if !has {
		common.ShowMsg("fail")
		return
	}
	model.Author=user.Id
	_,err=model.Add()
	if err!=nil {
		common.ShowMsg("fail")
		return
	}
	response.ShowData(c,model)
	return
}

func Edit(c *gin.Context)  {
	data,err:=request.GetJson(c)
	if err != nil {
		response.ShowError(c, "fail")
		return
	}
	if _, ok := data["id"]; !ok {
		response.ShowError(c, "fail")
		return
	}
	article:=models.SystemArticle{}
	article.Id = int64(data["id"].(float64))
	has:=article.GetRow()
	if !has {
		response.ShowError(c,"article_error")
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
	model :=models.SystemArticle{Id:article.Id}
	if !strings.Contains(data["image_uri"].(string),"http://"){
		model.ImageUri=common.Substr(data["image_uri"].(string),"upload/")
		//model.ImageUri,_=common.WriteFile("./upload",data["image_uri"].(string))
	}
	model.Title=data["title"].(string)
	model.Content=data["content"].(string)
	model.ContentShort=data["content_short"].(string)
	model.SourceUri=data["source_uri"].(string)
	//
	model.Importance=int(data["importance"].(float64))
	status:=int(data["status"].(float64))
	if status ==1{
		model.Status=1
	}
	if data["comment_disabled"].(bool) {
		model.CommentDisabled =1
	}
	model.Display_time = common.StrToTimes(data["display_time"].(string))

	user :=models.SystemUser{Name:data["author"].(string)}
	has =user.GetRow()
	if !has {
		common.ShowMsg("fail")
		return
	}
	model.Author=user.Id
	err=model.Update()
	if err!=nil {
		common.ShowMsg("fail")
		return
	}
	response.ShowData(c,model)
	return
}

func Index(c *gin.Context)  {
	page, _ := strconv.ParseInt(c.Query("page"), 10, 64)
	limit, _ := strconv.ParseInt(c.Query("limit"), 10, 64)
	type filters struct {
		Status int
		Title string
		Importance int
		StartTime time.Time
		EndTime time.Time
	}
    filter:=filters{}
	dateValues:=c.QueryArray("dateValue[]")
	if len(dateValues)==2 {
		filter.StartTime =  common.StrToTimes(dateValues[0])
		filter.EndTime =  common.StrToTimes(dateValues[1])
	}
	status :=c.Query("status")
	if status!="" {
		filter.Status,_=strconv.Atoi(status)
	}
	importance :=c.Query("importance")
	if importance!="" {
		filter.Importance,_=strconv.Atoi(importance)
	}
	filter.Title = c.Query("title")

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
func ShowImage(c *gin.Context){
	imgName := c.Query("imgname")
	//顾虑危险字符
	imgName = strings.Replace(imgName,"../","",-1)
	c.File(imgName)
}
func Detail(c *gin.Context){
	id,has:=c.GetQuery("id")
	if	!has{
		response.ShowErrorParams(c, "id")
		return
	}
	model:=models.SystemArticle{}
	model.Id, _ = strconv.ParseInt(id, 10, 64)
	has=model.GetRow()
	if !has {
		response.ShowError(c,"article_error")
		return
	}
	articles:=models.SystemArticlePage{}
	articles.SystemArticle=model
	userModel := models.SystemUser{}
	userModel.Id=model.Author
	has=userModel.GetRow()
	if has {
		articles.AuthorName=userModel.Name
	}
	if articles.ImageUri!="" {

		articles.ImageUri=conf.Cfg.Host+"/showimage?imgname=upload/"+articles.ImageUri
	}
	response.ShowData(c, articles)
	return
}