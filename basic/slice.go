package main

import "fmt"

func main() {

	var arraykita [3]string

	arraykita[0] = "zulfikar"

	println(arraykita[0])

	//make slice
	newSlice := make([]string, 7, 8)

	newSlice[0] = "ini adalah slice jika tidak dari array"

	//make slice form array

	newArray := [3]string{"ini", "adalalh", "array"}

	newSliceFromArray := newArray[1:2]

	fmt.Println(newSliceFromArray)

}
