package main

import (
	"fmt"
	"go-admin/models"
)

func main() {
	var constant []models.SystemMenu
	menu := models.SystemMenu{Type:1}
	constant,err:=menu.GetRowByType()
	fmt.Println(constant,err)

}