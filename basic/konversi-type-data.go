package main

import "fmt"

func main() {
	age := int16(200)
	age8 := int8(age)
	age = 20

	fmt.Println(age8, age)

	name := "Alexander Zull"
	name2 := name[2]

	//konversi dari byte ke string
	fmt.Println(string(name2))
}
