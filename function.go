package main

import (
	"fmt"
	"math/rand"
	"time"
)

// func main() {
// 	SayHello("")
// }
// func SayHello(name string) {
// 	name = "Zaki"
// 	fmt.Println("Hello ", name)
// }
// func main() {
// 	names := []string{"Ery", "Zakaria"}
// 	PrintMessage("Hello", names)
// }

// func PrintMessage(comment string, arr []string) {
// 	nameString := strings.Join(arr, " ")
// 	fmt.Println(comment, nameString)
// }

func main() {
	rand.Seed(time.Now().Unix())
	var randomRange int

	randomRange = randomWithValue(1, 100000)
	fmt.Println("Angka ke : ", randomRange)
}

func randomWithValue(min, max int) int {
	value := rand.Int()%(max-min+1) + min
	return value
}
