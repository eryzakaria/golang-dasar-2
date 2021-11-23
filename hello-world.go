package main

import "fmt"

func main() {
	/*
		Disini kita akan belajar
		bahasa pemrograman golang
	*/
	/*fmt.Print("Ini merupakan sebuah data \n") // Mirip dengan println, namun output tidak menghasilkan "enter"

	println("Hello, world")

	fmt.Println("Hello, world") // Baris ini mengarah ke import "fmt"

	fmt.Printf("Masukan Kata:\n")
	fmt.Printf("Masukan Kata:\n")
	fmt.Scanf("Merupakan sebuah kata, tuliskan:") // Untuk mendapatkan kata input*/

	// var cek string = "cek"
	// fmt.Println(cek)
	// cek = "kamu"

	// fmt.Println(cek)

	// one, two, three := "satu", "dua", "tiga"

	// four, five, six := 4, 5, 6

	// fmt.Println(one, two, three)
	// fmt.Println(four, five, six)

	// Biy, _ := "Joni", "rini"

	// fmt.Println(Biy)

	// Sumpah := new(int)
	// sumpah := *Sumpah

	// fmt.Println(&Sumpah)

	// var name = false

	// fmt.Printf("You Married Me %t", name)

	// FirstName := "Ery"
	// LastName := "Trimadhani"

	// fmt.Print("\n", FirstName, LastName, "\n")

	// fmt.Println(((20 * 8) / 2) % 3)

	// angka1 := 70
	// angka2 := 80

	// if angka1 != angka2 {
	// 	println("Nilai Tidak Sama")
	// } else {
	// 	println("Nilai Sama")
	// }

	// point := 6

	// switch {
	// case point == 8:
	// 	fmt.Println("Nilai Memuaskan")
	// 	fallthrough
	// case (point < 8) && (point > 5):
	// 	fmt.Println("Cukup")
	// 	fallthrough
	// default:
	// 	fmt.Println("Gak Tau")
	// }

	// nilai := 5

	// if nilai > 7 {
	// 	switch {
	// 	case nilai == 10:
	// 		fmt.Println("Keren")
	// 	default:
	// 		fmt.Println("Lumayan")
	// 	}
	// } else if nilai > 6 {
	// 	switch {
	// 	case nilai == 6:
	// 		fmt.Println("Lulus")
	// 	}
	// } else {
	// 	fmt.Println("Ngulang")
	// }

	// switch {
	// case nilai > 7:
	// 	if nilai == 10 {
	// 		fmt.Println("Keren")
	// 	} else {
	// 		fmt.Println("Lulus")
	// 	}
	// case nilai > 5:
	// 	fmt.Println("Kurang")
	// }

	// for nilai < 10 {
	// 	fmt.Println("Angka", nilai)
	// 	nilai++
	// }

	// i := 0

	// for {
	// 	fmt.Println("Angka", i)
	// 	i++
	// 	if i == 5 {
	// 		break
	// 	}
	// }

	// for kk := 1; kk <= 10; kk++ {
	// 	if kk%2 == 1 {
	// 		fmt.Println(kk)
	// 		continue

	// 	} else if kk > 8 {
	// 		break
	// 	}
	// 	fmt.Println("Angka", kk)
	// }

	// var input int

	// fmt.Printf("Masukan Angka: ")
	// fmt.Scanf("%d", &input)

	// for i := 1; i <= input; i++ {
	// 	for j := 1; j <= input; j++ {
	// 		fmt.Printf("%d\t", j*i)
	// 	}
	// 	fmt.Println()
	// }

	// outerloop:
	// 	for i := 0; i < 5; i++ {
	// 		for j := 0; j < 5; j++ {
	// 			if i == 5 {
	// 				break outerloop

	// 			}
	// 			fmt.Println("Matriks", "[", i, "]", "[", j, "]")
	// 		}
	// 	}

	// for z := 0; z < 5; z++ {
	// 	for a := 0; a < 5; a++ {
	// 		for b := 0; b < 5; b++ {
	// 			for c := 0; c < 5; c++ {
	// 				fmt.Printf("%d%d", c, b)
	// 			}
	// 			fmt.Println()
	// 		}
	// 	}
	// }

	// var Array [4]int
	// Array[0] = 1
	// Array[1] = 2
	// Array[2] = 3

	// fmt.Println(Array)

	// fruits := [3]string{"Melon", "Jeruk", "Mangga"}

	// for i := 0; i < len(fruits); i++ {
	// 	fmt.Printf("Data %d = %s\n", i, fruits[i])
	// }

	// fmt.Println(len(fruits))
	// fmt.Println(fruits)

	// number := [...]int{1, 2, 3}

	// fmt.Println(number)

	// jumlah := [2][4]int{[4]int{1, 2, 3}, [4]int{1, 2, 3}}
	// jumlah2 := [2][4]int{{1, 2, 3}, {1, 2, 3}}

	// fmt.Println(len(jumlah))
	// fmt.Println(jumlah)
	// fmt.Println(jumlah2)

	// var fruits = [3]string{"Mangga", "Melon", "Durian"}

	// for i, _ := range fruits {
	// 	fmt.Printf("Buah %d\n", i)
	// }

	// var array = []string{
	// 	"Monyet",
	// 	"Gorila",
	// 	"Singa",
	// }

	// var name = array

	// fmt.Println(name[1])

	// name := make([]string, 2)

	// fmt.Println()

	// slice := []string{"jeruk", "Mangga", "durian"}
	// slices := slice[0:2]

	// fmt.Println(slices)

	// var fruits = []string{"Apple", "Mango", "Pear", "Orange"}

	// afruits := fruits[0:4]
	// bfruits := fruits[1:4]

	// aafruits := afruits[1:3]
	// bbfruits := bfruits[0:1]

	// fmt.Println(afruits)
	// fmt.Println(bfruits)
	// fmt.Println(aafruits)
	// fmt.Println(bbfruits)

	// bbfruits[0] = "Banana"

	// fmt.Println(bbfruits)

	// month := []string{"Januari", "Februari", "Maret", "April", "Mei", "Juni", "Juli", "Agustus", "September", "Oktober", "November", "Desember"}
	// fmt.Println(cap(month))
	// months := month[4:7]

	// fmt.Println(len(months))
	// fmt.Println(cap(months))
	// fmt.Println(months)

	// months = append(months, "Desember")
	// fmt.Println(months)

	// fruits := []string{"Pepaya", "Jeruk", "Mangga", "Durian"}
	// fruit := fruits[1:3]
	// afruit := fruits[0:4:4]

	// fmt.Println(afruit)

	// fmt.Println(len(fruit))

	// fruitt := append(fruit, "Pisang")
	// fmt.Println(fruitt)
	// fmt.Println(fruit)
	// fmt.Println(fruits)

	// var fruits = []string{"apple", "grape", "banana"}
	// var bFruits = fruits[0:2]
	// fmt.Println(cap(bFruits)) // 3
	// fmt.Println(len(bFruits)) // 2
	// fmt.Println(fruits)       // ["apple", "grape", "banana"]
	// fmt.Println(bFruits)      // ["apple", "grape"]
	// var cFruits = append(bFruits, "papaya")
	// fmt.Println(fruits)  // ["apple", "grape", "papaya"]
	// fmt.Println(bFruits) // ["apple", "grape"]
	// fmt.Println(cFruits) // ["apple", "grape", "papaya"]

	// dst := make([]string, 3)
	// src := []string{"Melon", "Mangga", "Pisang", "Durian"}
	// n := copy(dst, src)

	// fmt.Println(dst)
	// fmt.Println(src)
	// fmt.Println(n)

	// var chicken map[string]int
	// chicken = map[string]int{}

	// chicken["Januari"] = 1
	// chicken["Februari"] = 2

	// fmt.Println("Januari", chicken["Januari"])
	// fmt.Println("Maret", chicken["Maret"])

	// var check map[int]int
	// check = map[int]int{}

	// check[2] = 1
	// check[3] = 2

	// fmt.Println("", check[2])
	// fmt.Println("Check", check[3])

	// ran := map[string]int{"Januari": 1, "Februari": 2}

	// fmt.Println(len(ran))
	// fmt.Println("Januari", ran["Januari"])
	// fmt.Println("Februari", ran["Februari"])

	// answer := map[int]string{1: "RAN", 2: "NOAH"}

	// fmt.Println("Qoute : ", answer[1])
	// fmt.Println("Qoute : ", answer[2])

	// tv := [3]string{
	// 	"Trans 7",
	// 	"RCTI",
	// 	"Global Tv",
	// }

	// for i, tvs := range tv {
	// 	fmt.Println(i, tvs)
	// }

	Document := []map[string]string{
		map[string]string{"Nama": "Ery Zakaria Trimadhani"},
		map[string]string{"Address": "Tegal"},
	}

	for i, Documents := range Document {
		fmt.Println("Data : ", i, Documents)
	}

	rice := map[string]string{
		"Name":     "Ery Zakaria Trimadhani",
		"Adddress": "Tegal",
	}
	rice["Title"] = "Programmer"

	fmt.Println(rice)

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}

		fmt.Println("Bilangan Ganjil", i)
	}

}
