package main

import "fmt"

func main() {
	none := 6

	switch none {
	case 8:
		fmt.Println("Perfect")
	case 7:
		fmt.Println("Awesome")
	default:
		fmt.Println("BAD")
	} 

	none = 1
	switch none {
	case 10:
		fmt.Println("Sempurna")
	case 9,8,7:
		fmt.Println("Lumayan")
	case 6,5:
		fmt.Println("Cukup")
	default:
		fmt.Println("Pancen Goblok")
	}

	none = 7

	switch {
	case none == 8:
		fmt.Println("Perfect")
	case (none <= 7) && (none > 3):
		fmt.Println("Awesome")
	default:
		fmt.Println("Cukup Goblok")
	}
}