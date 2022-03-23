package menu

import (
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go-admin/lib/request"
	"go-admin/lib/response"
	"go-admin/models/systemnewdb"
	"io/ioutil"
	"strconv"
	"time"
	"go-admin/lib/common"

)

type Role struct {
	Id          int           `form:"id" json:"id"`
	Name        string        `form:"name" json:"name"`
	Description string        `form:"description" json:"description"`
	State       int8          `form:"state" json:"state"`
	Routes      []interface{} `form:"routes" json:"routes"`
}

func List(c *gin.Context) {
	session := sessions.Default(c)
	v := session.Get(common.Conf.Token)
	if v == nil {
		response.ShowError(c, "fail")
		return
	}
	uid := session.Get(v)
	db := systemnewdb.GetDb()

	user := systemnewdb.SystemUser{ID: uid.(int)}
	db.First(&user)

	menu := systemnewdb.SystemMenu{}
	menuArr, err := menu.GetAll()
	if err != nil {
		response.ShowError(c, "fail")
		return
	}
	var menuMap = make(map[int][]systemnewdb.SystemMenu, 0)
	role := systemnewdb.SystemRole{}
	mrArr := role.GetRowMenu()
	for _, value := range menuArr {
		value.Hidden = 0
		menuMap[value.Pid] = append(menuMap[value.Pid], value)
	}
	jsonArr := TreeMenuNew(menuMap, 0, mrArr)
	response.ShowData(c, jsonArr)
	return
}

func Roles(c *gin.Context) {
	model := systemnewdb.SystemRole{}
	menu := systemnewdb.SystemMenu{}
	roleArr, err := model.GetAll()
	if err != nil {
		response.ShowError(c, "fail")
		return
	}
	var roleMenu []Role
	for _, value := range roleArr {
		r := Role{}
		r.Id = value.ID
		r.Name = value.Name
		r.State = value.State
		r.Description = value.Description
		menuArr := menu.GetRouteByRole(value.ID)
		if menuArr != nil {
			var menuMap = make(map[int][]systemnewdb.SystemMenu, 0)
			for _, value := range menuArr {
				value.Hidden = 0
				menuMap[value.Pid] = append(menuMap[value.Pid], value)
			}
			role := systemnewdb.SystemRole{}
			mrArr := role.GetRowMenu()
			jsonStr := TreeMenuNew(menuMap, 0, mrArr)
			if jsonStr != nil {
				r.Routes = jsonStr
			}
		}
		roleMenu = append(roleMenu, r)
	}
	response.ShowData(c, roleMenu)
	return
}

func Dashboard(c *gin.Context) {
	session := sessions.Default(c)
	v := session.Get(common.Conf.Token)
	if v == nil {
		response.ShowError(c, "fail")
		return
	}
	uid := session.Get(v)
	user := systemnewdb.SystemUser{ID: uid.(int)}
	_, _ = user.GetRow()
	menu := systemnewdb.SystemMenu{State: 1}
	var menuArr []systemnewdb.SystemMenu
	var err error
	if user.Name == "admin" {
		menuArr, err = menu.GetAll()
		if err != nil {
			response.ShowError(c, "fail")
			return
		}
	} else {
		menuArr = menu.GetRowByUid(uid)
		fmt.Println(menuArr)
	}
	var menuMap = make(map[int][]systemnewdb.SystemMenu, 0)
	for _, value := range menuArr {
		menuMap[value.Pid] = append(menuMap[value.Pid], value)
	}
	role := systemnewdb.SystemRole{}
	mrArr := role.GetRowMenu()
	jsonStr := TreeMenuNew(menuMap, 0, mrArr)
	response.ShowData(c, jsonStr)
	return
	//roleMenu:="{\"menuList\":[{\"create_time\":\"2018-03-1611:33:00\",\"menu_type\":\"M\",\"children\":[{\"create_time\":\"2018-03-1611:33:00\",\"menu_type\":\"C\",\"children\":[],\"parent_id\":1,\"menu_name\":\"用户管理\",\"icon\":\"#\",\"perms\":\"system:user:index\",\"order_num\":1,\"menu_id\":4,\"url\":\"/system/user\"},{\"create_time\":\"2018-12-2810:36:20\",\"menu_type\":\"M\",\"children\":[{\"create_time\":\"2018-12-2810:50:28\",\"menu_type\":\"C\",\"parent_id\":73,\"menu_name\":\"人员通讯录\",\"icon\":null,\"perms\":\"system:person:index\",\"order_num\":1,\"menu_id\":74,\"url\":\"/system/book/person\"}],\"parent_id\":1,\"menu_name\":\"通讯录管理\",\"icon\":\"fafa-address-book-o\",\"perms\":null,\"order_num\":1,\"menu_id\":73,\"url\":\"/system\"}],\"parent_id\":0,\"menu_name\":\"系统管理\",\"icon\":\"fafa-adjust\",\"perms\":null,\"order_num\":2,\"menu_id\":1,\"url\":\"#\"},{\"create_time\":\"2018-03-1611:33:00\",\"menu_type\":\"M\",\"children\":[{\"create_time\":\"2018-03-1611:33:00\",\"menu_type\":\"C\",\"parent_id\":2,\"menu_name\":\"数据监控\",\"icon\":\"#\",\"perms\":\"monitor:data:view\",\"order_num\":3,\"menu_id\":15,\"url\":\"/system/druid/monitor\"}],\"parent_id\":0,\"menu_name\":\"系统监控\",\"icon\":\"fafa-video-camera\",\"perms\":null,\"order_num\":5,\"menu_id\":2,\"url\":\"#\"}],\"user\":{\"login_name\":\"admin\",\"user_id\":1,\"user_name\":\"管理员\",\"dept_id\":1}}"
	////roleMenu:="{\"menuList\": [{ 		\"children\": [{ 			\"menu_type\": \"M\", 			\"children\": [{ 				\"menu_type\": \"C\", 				\"parent_id\": 73, 				\"menu_name\": \"人员通讯录\", 				\"icon\": null, 				\"order_num\": 1, 				\"menu_id\": 74, 				\"url\": \"/system/book/person\" 			}], 			\"parent_id\": 1, 			\"menu_name\": \"通讯录管理\", 			\"icon\": \"fafa-address-book-o\", 			\"perms\": null, 			\"order_num\": 1, 			\"menu_id\": 73, 			\"url\": \"#\" 		}], 		\"parent_id\": 0, 		\"menu_name\": \"系统管理\", 		\"icon\": \"fafa-adjust\", 		\"perms\": null, 		\"order_num\": 2, 		\"menu_id\": 1, 		\"url\": \"#\" 	}], 	\"user\": { 		\"login_name\": \"admin1\", 		\"user_id\": 1, 		\"user_name\": \"管理员\", 		\"dept_id\": 1 	} }"
	//var data map[string]interface{}
	//err := json.Unmarshal([]byte(roleMenu), &data)
	//if err != nil {
	//	fmt.Println(err)
	//}

	//c.JSON(http.StatusOK, gin.H{
	//	"code": 20000,
	//	"data":  data,
	//})
	//return
}
func TreeMenuNew(menuMap map[int][]systemnewdb.SystemMenu, pid int, mrArr map[int][]string) []interface{} {
	var menuNewArr []interface{}
	if _, ok := menuMap[pid]; ok {
		for _, value := range menuMap[pid] {
			var item = make(map[string]interface{})
			item["path"] = value.Path
			item["component"] = value.Component
			if value.Redirect != "" {
				item["redirect"] = value.Redirect
			}
			//if value.Alwaysshow == 1 {
			//	item["alwaysShow"] = true
			//}
			if value.Hidden == 1 {
				item["hidden"] = true
			} else {
				item["hidden"] = false
			}
			var meta = make(map[string]interface{})
			if _, ok := mrArr[value.ID]; ok {
				meta["roles"] = mrArr[value.ID]
			}
			if value.MetaTitle != "" {
				meta["title"] = value.MetaTitle
			}
			if value.MetaIcon != "" {
				meta["icon"] = value.MetaIcon
			}
			if value.MetaAffix == 1 {
				meta["affix"] = true
			}
			if value.MetaNocache == 1 {
				meta["noCache"] = true
			}
			if value.State == 1 {
				meta["status"] = true
			}

			if len(meta) > 0 {
				item["meta"] = meta
			}
			item["pid"] = value.Pid
			item["id"] = value.ID
			item["url"] = value.URL
			item["name"] = value.Name
			children := TreeMenuNew(menuMap, value.ID, mrArr)
			if children != nil {
				item["children"] = children
			}
			menuNewArr = append(menuNewArr, item)
		}
	}
	return menuNewArr
}

func Index(c *gin.Context) {
	menu := systemnewdb.SystemMenu{}
	menuArr, err := menu.GetAll()
	if err != nil {
		response.ShowError(c, "fail")
		return
	}
	fmt.Println(menuArr)
	var menuMap = make(map[int][]systemnewdb.SystemMenu, 0)
	for _, value := range menuArr {
		menuMap[value.Pid] = append(menuMap[value.Pid], value)
	}
	var menuNewArr []systemnewdb.SystemMenu
	menuNewArr = TreeNode(menuMap, 0)
	response.ShowData(c, menuNewArr)
	return
}
func TreeNode(menuMap map[int][]systemnewdb.SystemMenu, pid int) []systemnewdb.SystemMenu {
	var menuNewArr []systemnewdb.SystemMenu
	if _, ok := menuMap[pid]; ok {
		for _, v := range menuMap[pid] {
			menuNewArr = append(menuNewArr, v)
			menuNewArr = append(menuNewArr, TreeNode(menuMap, v.ID)...)
		}
	}
	return menuNewArr
}

func Create(c *gin.Context) {
	jsonstr, _ := ioutil.ReadAll(c.Request.Body)
	var data map[string]interface{}
	err := json.Unmarshal(jsonstr, &data)
	if err != nil {
		response.ShowError(c, "fail")
		return
	}
	if _, ok := data["name"]; !ok {
		response.ShowError(c, "fail")
		return
	}
	if _, ok := data["path"]; !ok {
		response.ShowError(c, "fail")
		return
	}
	if _, ok := data["component"]; !ok {
		response.ShowError(c, "fail")
		return
	}
	if _, ok := data["url"]; !ok {
		response.ShowError(c, "fail")
		return
	}

	menus := systemnewdb.SystemMenu{}
	menus.Path = data["path"].(string)
	if menus.Path == "" {
		response.ShowError(c, "fail")
		return
	}
	has := menus.GetRow()
	if has > 0 && menus.Path != "#" {
		response.ShowError(c, "path不可重复")
		return
	}
	menu := systemnewdb.SystemMenu{}
	menu.Name = data["name"].(string)
	menu.Path = data["path"].(string)
	menu.MetaTitle = data["name"].(string)
	menu.Component = data["component"].(string)
	menu.URL = data["url"].(string)
	menu.Redirect = data["redirect"].(string)
	menu.MetaIcon = data["meta_icon"].(string)
	if data["alwaysshow"].(bool) {
		menu.Alwaysshow = 1
	}
	if data["hidden"].(bool) {
		menu.Hidden = 1
	}
	if data["state"].(bool) {
		menu.State = 1
	}
	menu.Ctime = time.Now()

	menu.Sort, _ = strconv.Atoi(data["sort"].(string))
	menu.Pid = int(data["pid"].(float64))
	if menu.Pid == 0 {
		menu.Level = int8(menu.Pid)
	} else {
		pidMenuModel := systemnewdb.SystemMenu{ID: menu.Pid}
		_ = pidMenuModel.GetRow()
		menu.Level = pidMenuModel.Level + 1
	}
	db := systemnewdb.GetDb()
	result := db.Create(&menu)
	if result.Error != nil {
		response.ShowError(c, "fail")
		return
	}
	response.ShowData(c, menu)
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
	menu := systemnewdb.SystemMenu{}
	menu.ID = int(data["id"].(float64))
	has := menu.GetRow()
	if has > 0 {
		response.ShowError(c, "要修改数据不存在")
		return
	}
	if _, ok := data["name"]; !ok {
		response.ShowError(c, "fail")
		return
	}
	if _, ok := data["path"]; !ok {
		response.ShowError(c, "fail")
		return
	}
	if _, ok := data["component"]; !ok {
		response.ShowError(c, "fail")
		return
	}
	if _, ok := data["url"]; !ok {
		response.ShowError(c, "fail")
		return
	}
	menuModel := systemnewdb.SystemMenu{}
	menuModel.Path = data["path"].(string)
	if menuModel.Path == "" {
		response.ShowError(c, "fail")
		return
	}
	has = menuModel.GetRow()
	if has > 0 && menuModel.Path != "#" && menuModel.ID != menu.ID {
		response.ShowError(c, "path不可重复")
		return
	}
	menu.Name = data["name"].(string)
	menu.Path = data["path"].(string)
	menu.MetaTitle = data["name"].(string)
	menu.Component = data["component"].(string)
	menu.URL = data["url"].(string)
	menu.Redirect = data["redirect"].(string)
	menu.MetaIcon = data["meta_icon"].(string)
	if data["meta_nocache"].(bool) {
		menu.MetaNocache = 1
	}
	if data["hidden"].(bool) {
		menu.Hidden = 1
	}
	if data["alwaysshow"].(bool) {
		menu.Alwaysshow = 1
	}

	if data["state"].(bool) {
		menu.State = 1
	}
	menu.Sort, _ = strconv.Atoi(data["sort"].(string))
	menu.Pid = int(data["pid"].(float64))
	db := systemnewdb.GetDb()
	result := db.Save(&menu)
	if result.Error != nil {
		response.ShowError(c, "fail")
		return
	}
	response.ShowData(c, "success")
	return
}
func Delete(c *gin.Context) {
	str, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		response.ShowError(c, "fail")
		return
	}
	id := string(str)
	if id == "" {
		response.ShowError(c, "fail")
		return
	}
	menu := systemnewdb.SystemMenu{}
	menu.ID, _ = strconv.Atoi(id)

	err = menu.Delete()
	if err != nil {
		response.ShowError(c, "fail")
		return
	}
	response.ShowData(c, "success")
	return
}
