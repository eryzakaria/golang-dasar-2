package main

import "fmt"

func main() {
	result := sayHello("Ery", 1000000)

	fmt.Println(result)
}

func sayHello(Name string, Balance int) string {

	fmt.Println(Name, Balance)

	return "Hello "
}
