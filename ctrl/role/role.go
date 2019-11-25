package role

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-admin/models"
	"go-admin/modules/response"
	"io/ioutil"
)

func UpdateRole(c *gin.Context)  {
	jsonstr, _ := ioutil.ReadAll(c.Request.Body)
	var data map[string]interface{}
	err := json.Unmarshal(jsonstr, &data)
	if err != nil {
		response.ShowError(c, "fail")
		return
	}
	model:=models.SystemRole{Name:data["name"].(string)}
	has:=model.GetRow()
	if !has {
		response.ShowError(c, "role_error")
		return
	}
	model.AliasName=data["name"].(string)
	model.Description=data["description"].(string)

	err=model.Update(data["routes"].([]interface{}))
	if err!=nil {
		fmt.Println(err)
		response.ShowError(c, "fail")
		return
	}
	datas:=map[string]string{"status":"success"}
	response.ShowData(c,datas)
	return
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
	err=model.AddCommit(data["routes"].([]interface{}))
	if err!=nil {
		response.ShowError(c, "fail")
		return
	}
	datas:=map[string]int{"key":model.Id}
	response.ShowData(c,datas)
	return
}
func DeleteRole(c *gin.Context) {
	name := c.Param("name") //通过Param获取
	if name=="" {
		response.ShowError(c, "fail")
		return
	}
	role:=models.SystemRole{Name:name}
	has:=role.GetRow()
	if !has {
		response.ShowError(c, "fail")
		return
	}

	roles:=models.SystemRole{Id:role.Id}
	err:=roles.Delete()
	if err!=nil {
		response.ShowError(c, "fail")
		return
	}
	datas:=map[string]string{"status":"success"}
	response.ShowData(c,datas)
	return
}

