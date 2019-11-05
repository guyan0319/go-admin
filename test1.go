package main

import (
	"fmt"
	"strconv"
)

func main()  {
	fmt.Println(strconv.FormatBool(0 < 1)) // true
	fmt.Println(strconv.FormatBool(0 > 1)) // false

}