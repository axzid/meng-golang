package main

func main() {

	age := 15

	switch age {
	case 15:
		println("You are 15 years old")
	case 20:
		println("You are 20 years old")
	default:
		println("You are not 15 or 20 years old")
	}

	//switch di go bisa di lakukan tanpa variable
	switch {
	case age == 15:
		println("You are 15 years old")
	case age == 20:
		println("You are 20 years old")
	default:
		println("You are not 15 or 20 years old")
	}
}
