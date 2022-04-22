package role

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-admin/models"
	"go-admin/modules/request"
	"go-admin/modules/response"
	"io/ioutil"
	"time"
)

func UpdateRole(c *gin.Context)  {
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
	id:=int(data["id"].(float64))
	model:=models.SystemRole{Name:data["name"].(string)}
	has:=model.GetRow()
	if has && model.Id!=id {
		response.ShowError(c, "role_error")
		return
	}
	model.Id=id
	model.AliasName=data["name"].(string)
	model.Description=data["description"].(string)
	if _, ok := data["status"]; ok{
		if data["status"].(bool) {
			model.Status=1
		}
	}
	var ids []int
	if _, ok := data["routes"]; ok {
		ids=TreeRoutes(data["routes"].([]interface{}))
	}
	err=model.Update(ids)
	if err!=nil {
		response.ShowError(c, "fail")
		return
	}
	datas:=map[string]string{"status":"success"}
	response.ShowData(c,datas)
	return
}
func TreeRoutes(routes []interface{} ) []int{
	var ids []int
	for _,value:=range routes {
		ids = append(ids,int( value.(map[string]interface{})["id"].(float64)))
		if _, ok := value.(map[string]interface{})["children"]; ok {
			children:=value.(map[string]interface{})["children"].([]interface{})
			ids = append(ids,TreeRoutes(children)...)
		}
	}
	return ids

}
func AddRole(c *gin.Context)  {
	jsonstr, _ := ioutil.ReadAll(c.Request.Body)
	var data map[string]interface{}
	err := json.Unmarshal(jsonstr, &data)
	if err != nil {
		response.ShowError(c, "fail")
		return
	}
	model:=models.SystemRole{Name:data["name"].(string)}
	has:=model.GetRow()
	if has {
		response.ShowError(c, "fail")
		return
	}
	model.AliasName=data["name"].(string)
	model.Description=data["description"].(string)
	if _, ok := data["status"]; ok{
		if data["status"].(bool) {
			model.Status=1
		}
	}
	model.Ctime=time.Now()
	err=model.AddCommit(data["routes"].([]interface{}))
	if err!=nil {
		response.ShowError(c, "fail")
		return
	}
	fmt.Println(model)
	response.ShowData(c,model)
	return
}
func DeleteRole(c *gin.Context) {
	data,err:=request.GetJson(c)
	if err != nil {
		response.ShowError(c, "fail")
		return
	}
	if _, ok := data["id"]; !ok {
		response.ShowError(c, "fail")
		return
	}
	id:=int(data["id"].(float64))
	role:=models.SystemRole{Id:id}
	has:=role.GetRow()
	if !has {
		response.ShowError(c, "fail")
		return
	}
	roles:=models.SystemRole{Id:role.Id}
	err=roles.Delete()
	if err!=nil {
		response.ShowError(c, "fail")
		return
	}
	datas:=map[string]string{"status":"success"}
	response.ShowData(c,datas)
	return
}

func Index(c *gin.Context)  {
	roles:=models.SystemRole{}
	list :=roles.GetNameList()
	response.ShowData(c,list)
	return
}