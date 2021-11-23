package main

import "fmt"

// func main() {
// 	var s, k int

// 	fmt.Printf("Hitung Sisi Persegi : ")
// 	fmt.Scanln(&s)
// 	fmt.Printf("Hitung Keliling Persegi: ")
// 	fmt.Scanln(&k)

// 	l := s * s

// 	fmt.Println("Hasil Luas Persegi: ", l)
// 	l = k * s
// 	fmt.Println("Hasil Keliling Persegi: ", l)

// }

func main() {
	var l, s float64

	fmt.Printf("Masukan Luas Bangunan : ")
	fmt.Scanln(&l)
	fmt.Printf("Masukan Luas Persegi : ")
	fmt.Scanln(&s)
	luas := l * (s * s)

	fmt.Println("Hasil : ", luas)

}
