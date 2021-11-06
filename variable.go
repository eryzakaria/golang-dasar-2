package main

import "fmt"

func main() {
	var (
		firstname = "Ery"
		middlename = "Zakaria"
		lastname = "Trimadhani"
	)
	var age int = 22

	fmt.Println("Nama Saya",firstname,middlename,lastname,"Umur",age)

	fmt.Printf("Perkenalkan nama saya %s %s %s usia saya sekarang %d tahun\n", firstname, middlename, lastname, age)

	number := 14

	fmt.Println("Tanggal Lahir",number)

	var name = "zaki"
	var angka = 2

	fmt.Println(name)
	fmt.Println(angka)

	var satu, dua, tiga, empat, lima string
	satu = "satu"
	dua = "dua"
	tiga = "tiga"
	empat = "empat"
	lima = "lima"

	fmt.Println(satu, dua, tiga, empat, lima)
	var s, d, t string
	s, d, t = "satu","dua","tiga"

	fmt.Println(s, d, t)
}