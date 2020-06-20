package main

import (
	"fmt"
	"strings"
)

func main() {
	var strs ="hello 你好 hello world"
	fmt.Println(Substr(strs,"好"))
}

func Substr(s,substr string)string{
	n:=strings.Index(s,substr)
	return s[n:]
}
