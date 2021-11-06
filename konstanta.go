package main

import "fmt"

func main() {
	const name string = "Ery"
	// name = "zaki" -> error, karena data string tidak dapat dirubah

	fmt.Printf("Nama saya %s", name)
}