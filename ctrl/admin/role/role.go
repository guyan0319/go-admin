package role

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"time"
	"go-admin/lib/request"
	"go-admin/lib/response"
	"go-admin/models/systemnewdb"
)

func UpdateRole(c *gin.Context) {
	jsonstr, _ := ioutil.ReadAll(c.Request.Body)
	var data map[string]interface{}
	err := json.Unmarshal(jsonstr, &data)
	if err != nil {
		response.ShowError(c, "fail")
		return
	}
	if _, ok := data["id"]; !ok {
		response.ShowError(c, "fail")
		return
	}
	id := int(data["id"].(float64))
	model := systemnewdb.SystemRole{Name: data["name"].(string)}
	has, _ := model.GetRow()
	if has > 0 && model.ID != id {
		response.ShowError(c, "role_error")
		return
	}
	model.ID = id
	model.AliasName = data["name"].(string)
	model.Description = data["description"].(string)
	if _, ok := data["state"]; ok {
		if data["state"].(bool) {
			model.State = 1
		}
	}
	var ids []int
	if _, ok := data["routes"]; ok {
		ids = TreeRoutes(data["routes"].([]interface{}))
	}
	err = model.Update(ids)
	if err != nil {
		response.ShowError(c, "fail")
		return
	}
	datas := map[string]string{"status": "success"}
	response.ShowData(c, datas)
	return
}
func TreeRoutes(routes []interface{}) []int {
	var ids []int
	for _, value := range routes {
		ids = append(ids, int(value.(map[string]interface{})["id"].(float64)))
		if _, ok := value.(map[string]interface{})["children"]; ok {
			children := value.(map[string]interface{})["children"].([]interface{})
			ids = append(ids, TreeRoutes(children)...)
		}
	}
	return ids

}
func AddRole(c *gin.Context) {
	jsonstr, _ := ioutil.ReadAll(c.Request.Body)
	var data map[string]interface{}
	err := json.Unmarshal(jsonstr, &data)
	if err != nil {
		response.ShowError(c, "fail")
		return
	}
	model := systemnewdb.SystemRole{Name: data["name"].(string)}
	has, _ := model.GetRow()
	if has > 0 {
		response.ShowError(c, "fail")
		return
	}
	model.AliasName = data["name"].(string)
	model.Description = data["description"].(string)
	if _, ok := data["state"]; ok {
		if data["state"].(bool) {
			model.State = 1
		}
	}
	model.Ctime = time.Now()
	err = model.AddCommit(data["routes"].([]interface{}))
	if err != nil {
		response.ShowError(c, "fail")
		return
	}
	response.ShowData(c, model)
	return
}
func DeleteRole(c *gin.Context) {
	data, err := request.GetJson(c)
	if err != nil {
		response.ShowError(c, "fail")
		return
	}
	if _, ok := data["id"]; !ok {
		response.ShowError(c, "fail")
		return
	}
	id := int(data["id"].(float64))
	role := systemnewdb.SystemRole{ID: id}
	has, _ := role.GetRow()
	if has < 1 {
		response.ShowError(c, "fail")
		return
	}
	roles := systemnewdb.SystemRole{ID: role.ID}
	err = roles.Delete()
	if err != nil {
		response.ShowError(c, "fail")
		return
	}
	datas := map[string]string{"status": "success"}
	response.ShowData(c, datas)
	return
}

func Index(c *gin.Context) {
	roles := systemnewdb.SystemRole{}
	list := roles.GetNameList()
	response.ShowData(c, list)
	return
}
