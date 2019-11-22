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
	model:=models.SystemRole{Name:data["key"].(string)}
	has:=model.GetRow()
	fmt.Println(has)
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










}