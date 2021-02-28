package main

import "fmt"

func main() {
	cityMap := map[string]string{
		"01": "BeiJing",
		"02": "ShangHai"}
	cityMap["03"] = "广州"
	delete(cityMap, "01")
	for key, item := range cityMap {
		fmt.Println(key, item)
	}

}
