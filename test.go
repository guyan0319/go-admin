package main

import "fmt"

func main() {
		var t=make(map[string]string)
		var t2=make(map[string]string,0)
		t["a"]="b"
		t2["a"]="b"

		fmt.Println(t,t2)
}