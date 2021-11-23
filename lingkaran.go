package main

import (
	"fmt"
)

// func main() {
// 	phi := 3.14
// 	var radius, luas float64

// 	fmt.Printf("Masukan Nilai jari-jari : ")
// 	fmt.Scanln(&radius)

// 	luas = phi * radius * radius

// 	fmt.Println(luas)

// }

// func main() {
// 	fmt.Println("Mencari Keliling Lingkaran")

// 	phi := 3.14
// 	var r float64
// 	fmt.Printf("Jari-jari : ")
// 	fmt.Scanln(&r)

// 	k := 2 * phi * r

// 	fmt.Println("Hasil", k)

// }

// func main() {
// 	var r, l float64
// 	var p float64

// 	fmt.Printf("Masukan Pecahan : ")
// 	fmt.Scanln(&p)
// 	fmt.Printf("Masukan Jari-Jari : ")
// 	fmt.Scanln(&r)

// 	l = p * 3.14 * r * r

// 	fmt.Println("Hasil", l)
// }

// func main() {
// 	var luas float64

// 	fmt.Printf("Masukan Luas Lingkaran : ")
// 	fmt.Scanln(&luas)

// 	r := luas / 3.14

// 	fmt.Println("Akar : ", r)
// 	fmt.Println(math.Sqrt(r))
// }

func main() {
	var kb, r float64

	fmt.Printf("Masukan Nilai Keliling Bangun : ")
	fmt.Scanln(&kb)
	fmt.Printf("Masukan Nilai Keliling Bangun : ")
	fmt.Scanln(&r)

	K1 := kb * 2 * 3.14 * r

	fmt.Println("Hasil : ", K1)

}
