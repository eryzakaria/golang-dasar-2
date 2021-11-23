package main

import "fmt"

func main() {
	// sumof3 := 0
	// sumof5 := 0

	fmt.Println("Rumus Barisan Bilangan")
	var Un, a, b, n int

	fmt.Printf("Masukan Bilangan Pertama : ")
	fmt.Scanln(&a)

	fmt.Printf("masukan beda angka : ")
	fmt.Scanln(&b)

	fmt.Printf("masukan banyak angka ke-n : ")
	fmt.Scanln(&n)

	Un = a + (n - 1)
	fmt.Println("Jumlah Suku ", Un)

}
