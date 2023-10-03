package main

import "fmt"

func lawakGemingClub(jumlah int) string {
	//convert int to string in golang
	return fmt.Sprintf("Jumlah lawak geming club adalah %d", jumlah)

}

// function return multiple value
func namaUmur(nama string, umur int) (string, int) {
	return fmt.Sprintf("nama : %s", nama), umur
}

func main() {

	fmt.Println(lawakGemingClub(10))

	// function return multiple value
	nama, umur := namaUmur("Rizky", 20)

	fmt.Println(nama)
	fmt.Println("umur:", umur)

}
