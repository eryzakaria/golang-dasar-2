package main

import "fmt"

func main() {
	var jumlah int

	fmt.Printf("Masukan Angka : ")
	fmt.Scanln(&jumlah)

	for i := 1; i < jumlah; i++ {
		if i%2 == 0 {
			fmt.Println("Genap")
		} else {
			fmt.Println("Ganjil")
		}
	}
}
