package main

import (
	"go-admin/models"
	"time"
)

func main() {
		r:=models.SystemRole{Name:"admin",Ctime:time.Now()}
		r.Add()
}