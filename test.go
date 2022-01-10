package main

import (
	"fmt"
	"math"
)

// func main() {
// 	result := sayHello("Ery", 1000000)

// 	fmt.Println(result)
// }

// func sayHello(Name string, Balance int) string {

// 	fmt.Println(Name, Balance)

// 	return "Hello "
// }

// func main() {
// 	sayHello()
// }

// func sayHello() (string string) {
// 	return "Ery" + "Zakaria"
// }
// func main() {
// 	name := mainMenu("Zaki")
// 	fmt.Println(name)

// 	fmt.Println(mainMenu(""))
// }

// func mainMenu(Name string) string {
// 	if Name == "" {
// 		return "Hallo Kenalan Dong"
// 	} else {
// 		return "Hello " + Name
// 	}
// }

// func main() {
// 	_, lastName := mainMenu()

// 	fmt.Println("Hello ", lastName)
// }

// func mainMenu() (string, string) {
// 	return "Ery", "Zakaria"
// }

// func main() {
// 	firstName, total := sayNumber(10, 10, 10, 10, 10)

// 	fmt.Println("Hallo", firstName, total)

// }

// func sayNumber(number ...int) (string, int) {
// 	total := 0
// 	for _, numbers := range number {
// 		total += numbers
// 	}
// 	return "zaki", total
// }

// func getFullName() (firstName, middleName, lastName string) {
// 	firstName = "Ery"
// 	middleName = "Zakaria"
// 	lastName = "Trimadhani"
// 	return
// }

// func main() {
// 	a, b, c := getFullName()
// 	fmt.Println(a, b, c)
// 	fmt.Println(a)
// 	fmt.Println(b)
// 	fmt.Println(c)
// }

// func sumAll(number ...int) int {
// 	a := 0
// 	for _, numbers := range number {
// 		a = a + numbers
// 	}
// 	return a
// }

// func main() {
// 	a := sumAll(10, 10, 10, 10, 10)
// 	fmt.Println(a)

// 	numbers := []int{10, 20, 30, 40, 50}
// 	a = sumAll(numbers...)
// 	fmt.Println(a)
// }

// func goodBye(name string) string {
// 	name = "Ery"
// 	return "Hello " + name
// }

// func main() {
// 	fmt.Println(goodBye(" "))
// }

// func goodGame(fullName string) string {
// 	return "Hello " + fullName
// }

// func main() {
// 	goodGame := goodGame("Ery")
// 	fmt.Println(goodGame)
// }

// func Lingkaran(phi float64) float64 {
// 	return 2 * phi
// }

// func main() {
// 	cek := Lingkaran(3.14)

// 	// var r float64
// 	// fmt.Printf("Masukan Data : ")
// 	// fmt.Scanln(&r)

// 	// hasil := cek * r
// 	// fmt.Println("Ini hasilnya : ", hasil)
// 	var r float64
// 	phi := 3.14
// 	fmt.Printf("Masukan jari-jari : ")
// 	fmt.Scanln(&r)
// 	hasil := cek * phi * r
// 	fmt.Println("Hasilnya ", hasil)
// }

// func main() {
// 	fmt.Println(strings.Trim("a  Ery Zakaria Trimadhani   a", "a"))
// }

func main() {
	fmt.Println(math.Max(10, 8))
}
