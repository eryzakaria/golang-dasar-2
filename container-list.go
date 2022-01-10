package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()

	data.PushFront("Ery")
	data.PushBack("Zakaria")
	data.PushBack("Trimadhani")

	for element := data.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}
}
