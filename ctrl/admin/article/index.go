package article

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	"web-demo/lib/common"
	"web-demo/lib/request"
	"web-demo/lib/response"
	"web-demo/models/systemdb"
)

func Create(c *gin.Context) {
	data, err := request.GetJson(c)
	if err != nil {
		response.ShowError(c, "fail")
		return
	}

	if _, ok := data["state"]; !ok {
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

	if _, ok := data["contentShort"]; !ok {
		response.ShowError(c, "fail")
		return
	}

	if _, ok := data["sourceUri"]; !ok {
		response.ShowError(c, "fail")
		return
	}
	if _, ok := data["imageUri"]; !ok {
		response.ShowError(c, "fail")
		return
	}
	if _, ok := data["displayTime"]; !ok {
		response.ShowError(c, "fail")
		return
	}
	if _, ok := data["commentDisabled"]; !ok {
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
	if data["title"].(string) == "" {
		response.ShowErrorParams(c, "title")
		return
	}
	if data["imageUri"].(string) == "" {
		response.ShowErrorParams(c, "imageUri")
		return
	}
	//fmt.Println(data)
	model := systemdb.SystemArticle{}
	//model.ImageUri,_=common.WriteFile("./upload",data["imageUri"].(string))
	model.ImageURI = common.SubstrPos(data["imageUri"].(string), "upload/")
	model.Title = data["title"].(string)
	model.Content = data["content"].(string)
	//model.Content=common.Base64Content(conf.Cfg.Host+"/showimage?imgname=upload/","./upload",model.Content)
	model.ContentShort = data["contentShort"].(string)
	model.SourceURI = data["sourceUri"].(string)

	model.Importance = uint8(data["importance"].(float64))
	status := int(data["state"].(float64))
	if status == 1 {
		model.State = 1
	}
	if data["commentDisabled"].(bool) {
		model.CommentDisabled = 1
	}
	model.DisplayTime = common.StrToTimes(data["displayTime"].(string))
	model.Ctime = int(model.DisplayTime.Unix())
	model.Mtime = model.DisplayTime
	user := systemdb.SystemUser{Name: data["author"].(string)}
	has, _ := user.GetRow()
	if has < 1 {
		common.ShowMsg("fail")
		return
	}
	model.Author = user.ID
	db := systemdb.GetDb()
	res := db.Create(&model)
	if res.Error != nil {
		fmt.Println(res.Error)
		common.ShowMsg("fail")
		return
	}
	response.ShowData(c, model)
	return
}

func Edit(c *gin.Context) {
	data, err := request.GetJson(c)
	if err != nil {
		response.ShowError(c, "fail")
		return
	}
	if _, ok := data["id"]; !ok {
		response.ShowError(c, "fail")
		return
	}
	article := systemdb.SystemArticle{}
	article.ID = uint64(data["id"].(float64))
	has, _ := article.GetRow()
	if has < 1 {
		response.ShowError(c, "article_error")
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
	if _, ok := data["contentShort"]; !ok {
		response.ShowError(c, "fail")
		return
	}
	if _, ok := data["sourceUri"]; !ok {
		response.ShowError(c, "fail")
		return
	}
	if _, ok := data["imageUri"]; !ok {
		response.ShowError(c, "fail")
		return
	}
	if _, ok := data["displayTime"]; !ok {
		response.ShowError(c, "fail")
		return
	}
	if _, ok := data["commentDisabled"]; !ok {
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
	if data["title"].(string) == "" {
		response.ShowErrorParams(c, "title")
		return
	}
	if data["imageUri"].(string) == "" {
		response.ShowErrorParams(c, "imageUri")
		return
	}
	model := systemdb.SystemArticle{ID: article.ID}
	if !strings.Contains(data["imageUri"].(string), "http://") {
		model.ImageURI = common.SubstrPos(data["imageUri"].(string), "upload/")
		//model.ImageUri,_=common.WriteFile("./upload",data["imageUri"].(string))
	}
	model.Title = data["title"].(string)
	model.Content = data["content"].(string)
	model.ContentShort = data["contentShort"].(string)
	model.SourceURI = data["sourceUri"].(string)
	//
	model.Importance = uint8(data["importance"].(float64))
	status := int(data["state"].(float64))
	if status == 1 {
		model.State = 1
	}
	if data["commentDisabled"].(bool) {
		model.CommentDisabled = 1
	}
	model.DisplayTime = common.StrToTimes(data["displayTime"].(string))

	user := systemdb.SystemUser{Name: data["author"].(string)}
	has, _ = user.GetRow()
	if has < 1 {
		common.ShowMsg("fail")
		return
	}
	model.Author = user.ID
	err = model.Update()
	if err != nil {
		common.ShowMsg("fail")
		return
	}
	response.ShowData(c, model)
	return
}

func Index(c *gin.Context) {
	page, _ := strconv.ParseInt(c.Query("page"), 10, 64)
	limit, _ := strconv.ParseInt(c.Query("limit"), 10, 64)
	var filters = map[string]string{"state": "", "title": "", "importance": "", "start_time": "", "end_time": ""}

	dateValues := c.QueryArray("dateValue[]")
	if len(dateValues) == 2 {
		filters["start_time"] = dateValues[0]
		filters["end_time"] = dateValues[1]
	}
	filters["state"] = c.Query("state")
	filters["importance"] = c.Query("importance")
	filters["title"] = c.Query("title")
	paging := &common.Paging{Page: page, PageSize: limit}
	articleModel := systemdb.SystemArticle{}
	articleArr, err := articleModel.GetAllPage(paging, filters)
	var articlePageArr []systemdb.SystemArticlePage

	for _, v := range articleArr {
		userModel := systemdb.SystemUser{}
		userModel.ID = v.Author
		has, _ := userModel.GetRow()
		if has < 1 {
			continue
		}
		articlePageArr = append(articlePageArr, systemdb.SystemArticlePage{SystemArticle: v, AuthorName: userModel.Name})
	}
	if err != nil {
		response.ShowError(c, "fail")
		return
	}
	data := make(map[string]interface{})
	data["items"] = articlePageArr
	data["total"] = paging.Total
	response.ShowData(c, data)
	return
}
func ShowImage(c *gin.Context) {
	imgName := c.Query("imgname")
	//顾虑危险字符
	imgName = strings.Replace(imgName, "../", "", -1)
	fmt.Println(imgName)
	c.File(imgName)
}
func Detail(c *gin.Context) {
	id, has := c.GetQuery("id")
	if !has {
		response.ShowErrorParams(c, "id")
		return
	}
	model := systemdb.SystemArticle{}
	idint64, _ := strconv.ParseInt(id, 10, 64)
	model.ID = uint64(idint64)
	res, _ := model.GetRow()
	if res < 1 {
		response.ShowError(c, "article_error")
		return
	}
	articles := systemdb.SystemArticlePage{}
	articles.SystemArticle = model
	userModel := systemdb.SystemUser{}
	userModel.ID = model.Author
	r, _ := userModel.GetRow()
	if r > 0 {
		articles.AuthorName = userModel.Name
	}
	if articles.ImageURI != "" {
		articles.ImageURI = common.Url() + "/showimage?imgname=upload/" + articles.ImageURI
		//fmt.Println(articles.ImageURI)
	}
	//fmt.Println(articles.DisplayTime)
	//articles.DisplayTime, _ = time.Parse("2006-01-02 15:04:05", articles.DisplayTime.Format("2006-01-02 15:04:05"))
	//fmt.Println(articles.DisplayTime)
	response.ShowData(c, articles)
	return
}
