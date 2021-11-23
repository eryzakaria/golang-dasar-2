package main

import "fmt"

func main() {

	number := 0

	for number < 5 {
		fmt.Println("Cetak", number)
		number++
	}

	for i := 0; i <= 5; i++ {
		fmt.Println("Angka", i)
	}

	for i := 0; i < 5; i++ {
		fmt.Print(i)
	}
	fmt.Println()

	// for i := 0; i < 5; i++ {
	// 	for j := i; j < 5; j++ {
	// 		fmt.Print(j," ")
	// 	}
	// 	fmt.Println()
	// }

	// for i := 5; i > 0; i-- {
	// 	for j := i; j > 0; j-- {
	// 		fmt.Print(j," ")
	// 	}
	// 	fmt.Println()
	// }

	// for i := 5; i > 0; i-- {
	// 	for j := i; j > 0; j-- {
	// 		fmt.Print(j," ")
	// 	}
	// 	fmt.Println()
	// }

	// for i := 1; i <= 5; i++ {
	// 	for j := 4; j >= i; j-- {
	// 		fmt.Print("")
	// 	}
	// 	for k := i; k >= 1; k-- {
	// 		fmt.Print(k)
	// 	}
	// 	fmt.Println("")
	// }

	// for i := 1; i <= 5; i++ {
	// 	for j := 4; j >= i; j-- {
	// 		fmt.Print(" ")
	// 	}
	// 	for k := i; k >= 1; k-- {
	// 		fmt.Print(k)
	// 	}
	// 	fmt.Println("")
	// }

	// number = 0

	// for {
	// 	fmt.Println("Angka", number)

	// 	number++
	// 	if number == 5 {
	// 		break
	// 	}
	// }

}
