package main

import "fmt"

func main() {

	var newMap map[string]string = make(map[string]string)

	newMap["key"] = "value"
	newMap["key2"] = "value2"
 
	fmt.Println(newMap)

	newMap2 := map[string]string{
		"ley":  "loy",
		"apko": "ad",
	}

	fmt.Println(newMap2)

}
