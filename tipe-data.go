package main

import "fmt"

func main() {
	/*
	Belajar Tipe Data String
	Trdapat: string, int, float, boolean
	*/
	
	// Ini merupakan data Integer
	var positifnumber uint16 = 25760
	// var negatifnumber int16 = -25760
	var negatifnumber = -1251212

	fmt.Printf("Bilangan Positif %d \n",positifnumber)
	fmt.Printf("Bilangan Negatif %d \n",negatifnumber)

	/*
	Ini merupakan data float
	Float sendiri untuk menentukan nilai koma
	*/

	number := 8.50

	fmt.Printf("Bilangan Decimal %f \n", number)
	fmt.Printf("Bilangan Decimal %.3f \n", number) // memformat data numerik desimal menjadi string

	// Ini merupakan data Boolean
	exist := true

	fmt.Printf("exist: %t \n", exist)

	// Ini merupakan data String
	message := "Hallo, saya zaki"

	fmt.Printf("Thankyou, %s \n", message)

	// Menggunakan Tanda Betik(``)
	text := `Hallo, perkenalkan
	Nama saya: Ery Zakaria trimadhani
	Umur saya: 22 tahun`

	fmt.Println(text)

}