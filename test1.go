package main

import (
	"encoding/json"
	"fmt"
)

func main()  {
	var data map[string]interface{}
	jsonStr := "{\"Name\":\"test\",\"Age\":19,\"infos\":[{\"info1\":\"hello\"},{\"info2\":\"hello\"}],\"pic\":{\"pic1\":\"img.jpg\"}}"
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
	for _,value:=range data{
		switch value.(type) {
		case float64:
			fmt.Println(value)
		case string:
			fmt.Println(value)
		case []interface {}:
			for k,v:=range value.([]interface {}){
				fmt.Println(k,v)
			}
		case map[string]interface {}:
			for k,v:=range value.(map[string]interface {}){
				fmt.Println(k,v)
			}
		}
	}

}