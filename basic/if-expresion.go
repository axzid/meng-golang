package main

func main() {

	name := "zlu"

	if name == "zlu" {
		println("Hello zlu")
	} else if name == "ad" {
		println("Hello ad")
	} else {
		println("Hello stranger")
	}

	//if expresion pada golang juga suport short statment dimana ini adalah kondisi sebelum if

	if length := len(name); length > 5 {
		println("Your name is long")
	} else {
		println("Your name is short")
	}

}
