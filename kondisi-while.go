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

	none = 8

	switch {
	case none == 8:
		fmt.Println("Perfect")
		fallthrough
	case (none <= 7) && (none > 3):
		fmt.Println("Awesome")
		fallthrough
	default:
		fmt.Println("Cukup Goblok")
	}

	none = 3

	if none >= 7 {
		switch none {
		case 10:
			fmt.Println("Perfect")
		default:
			fmt.Println("Nice")
		}	
	} else {
		if none >= 6 || none >= 5{
			fmt.Println("Kurang")
		} else if none == 4 || none == 3{
			fmt.Println("Yakin Kurang")
		} else{
			fmt.Println("Goblok!")
		}
	}
	
}