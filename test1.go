package main

import "fmt"

func main()  {
	var i float32
	var j float64

	i=1.5
	j=1.5
	fmt.Println(int(i))   // 1
	fmt.Println(int64(i)) // 1
	fmt.Println(int32(i)) // 1
	fmt.Println(int(j))   // 1
	fmt.Println(int64(j)) // 1
	fmt.Println(int32(j)) // 1
}