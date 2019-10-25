package user

import (
	"github.com/gin-gonic/gin"
	"go-admin/models"
	"go-admin/modules/response"
	"go-admin/public/common"
)

func Login(c *gin.Context){
	nickname :=c.Query("nickname")
	passwd :=c.Query("passwd")
	if nickname=="" || passwd=="" {
		response.ShowError(c,"fail")
		return
	}
	user:=models.SystemUser{Nickname:nickname}
	has := user.GetRowByNickname()
	if !has {
		response.ShowError(c,"fail")
		return
	}
	if common.Sha1En(passwd+user.Salt)!= user.Password {
		response.ShowError(c,"fail")
		return
	}
	response.ShowSuccess(c,"success")
	return

}