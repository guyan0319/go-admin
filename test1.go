package main

import (
	"fmt"
	"go-admin/public/common"
	"strings"
)

func main() {
	paging:=&common.Paging{Page:1,PageSize:10}
	paging.Total=100
	paging.GetPages()
	fmt.Println(paging)
}

func Substr(s,substr string)string{
	n:=strings.Index(s,substr)
	return s[n:]
}
