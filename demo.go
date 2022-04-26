package main

import (
	"go-admin/lib/common"
)



func main() {
	type A struct {
		Name string
		Id int
	}
	type B struct {
		Name string
		Id int
	}
	a:=&A{}
	b:=B{}
	common.CopyStruct(a, b)
}