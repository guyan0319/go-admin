package main

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	type  ss struct {
		a string
	}
	type sss struct {
		ss
		b string
	}
	aa:=sss{}
	aa.a="afsdf"
	//bb:=ss{a:"fasf"}
	//aa.ss=bb
	fmt.Println(aa)
}

func GetCurrentPath() (string, error) {
	//file, err := exec.LookPath(os.Args[0])
	//if err != nil {
	//	return "", err
	//}
	path, err := filepath.Abs("")
	fmt.Println(path)
	if err != nil {
		return "", err
	}
	i := strings.LastIndex(path, "/")
	if i < 0 {
		i = strings.LastIndex(path, "\\")
	}
	if i < 0 {
		return "", errors.New(`error: Can't find "/" or "\".`)
	}
	return string(path[0 : i+1]), nil
}