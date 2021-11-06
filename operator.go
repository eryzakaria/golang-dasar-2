package main

import "fmt"

func main() {
	/*
	Nilai Operator Aritmatika
	(+ - * / %)
	*/
	
	fmt.Println("Hasil",21 * 10)

	value := (((2 + 6) % 3) * 4 - 2) / 3
	hasil := ( value == 2)

	fmt.Printf("Hasil dari value : %d Menghasilkan nilai (%t) \n", value, hasil)

	/*
	Nilai Perbandingan
	*/

	left := false
	right := true

	var convert = left && right

	fmt.Printf("Maka menghasilkan Boolean %t \n", convert)

	convert = left || right
	fmt.Printf("Maka menghasilkan Boolean %t \n", convert)

	convert = !left
	fmt.Printf("Hasil Kebalikan %t \n", convert)
}