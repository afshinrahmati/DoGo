package main

import (
	"container/list"
	"fmt"
)

func main() {
	var name string = "afshin"
	age := 45
	age = 90
	text(age, name)
}
func text(age int, name string) {
	fmt.Print(age, name)
	list.New().PushFront(41)
}
