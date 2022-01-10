package main

import "fmt"

func main() {
	// // sumof3 := 0
	// // sumof5 := 0

	// fmt.Println("Rumus Barisan Bilangan")
	// var Un, a, b, n int

	// fmt.Printf("Masukan Bilangan Pertama : ")
	// fmt.Scanln(&a)

	// fmt.Printf("masukan beda angka : ")
	// fmt.Scanln(&b)

	// fmt.Printf("masukan banyak angka ke-n : ")
	// fmt.Scanln(&n)

	// Un = a + (n - 1)
	// fmt.Println("Jumlah Suku ", Un)
	// firstName, lastName := Value(" ", " ")
	// fmt.Println("Hello ", firstName, lastName)

	// total := sayHello(10, 20, 30, 40, 50)
	// fmt.Println("Hasil", total)

	// slice := []int{10, 30, 50, 70, 90}
	// total = sayHello(slice...)
	// fmt.Println(total)
	// sayWithFilter("Ery", spamFilter)
	// filter := spamFilter

	// sayWithFilter("Anjing", filter)

	name := func(name string) bool {
		return name == "admin"
	}

	registerUser("ery", name)

}

// func Value(Address, Country string) (firstName, lastName string) {
// 	firstName = "Ery "
// 	lastName = "Trimadhani "
// 	Address = "Tegal City"
// 	Country = " Indonesia"

// 	return firstName, lastName + Address + Country
// }

// func sayHello(number ...int) int {
// 	total := 0
// 	for _, numbers := range number {
// 		total += numbers
// 	}
// 	return total
// }

// func sayWithFilter(name string, Filter func(string) string) {
// 	fmt.Println("Hello ", Filter(name))
// }

// func spamFilter(name string) string {
// 	if name == "Anjing" {
// 		return "...."
// 	} else {
// 		return name
// 	}
// }

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You Are Blocked", name)
	} else {
		fmt.Println("Welcome To ", name)
	}
}
