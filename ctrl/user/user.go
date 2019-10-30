package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
)


func Reg(c *gin.Context){
	
	param3,_ :=c.GetPostForm("username")
	fmt.Println(c.PostForm("password"))
	fmt.Println(param3)
	//nickname :=c.Query("nickname")
	//passwd :=c.Query("passwd")
	//fmt.Println(nickname)
	//if nickname=="" || passwd=="" {
	//	response.ShowError(c,"fail")
	//	return
	//}
	//salt :=common.GetRandomBoth(4)
	//passwd = common.Sha1En(passwd+salt)
	//fmt.Println(salt)
	//fmt.Println(passwd)


}