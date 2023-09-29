package main

import "fmt"

func main() {

	array := []string{"m", "alexander", "zulfikar"}

	for i := 0; i < len(array); i++ {
		fmt.Println(array[i])
	}

	mapas := map[string]string{
		"nama": "alexander",
		"umur": "20",
		"asal": "jakarta",
	}

	for key, value := range mapas {
		fmt.Println(key, value)
	}

}
