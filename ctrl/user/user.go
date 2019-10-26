package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-admin/modules/response"
	"go-admin/public/common"
)

func Reg(c *gin.Context){
	nickname :=c.Query("nickname")
	passwd :=c.Query("passwd")
	fmt.Println(nickname)
	if nickname=="" || passwd=="" {
		response.ShowError(c,"fail")
		return
	}
	salt :=common.GetRandomBoth(4)
	passwd = common.Sha1En(passwd+salt)
	fmt.Println(salt)
	fmt.Println(passwd)


}