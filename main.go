package main

import (
	"container/list"
	"fmt"

	mydate "myApp/data"
)

func main() {
	var name string = "afshin"
	age := 45
	age = 90

	text(age, name, lastName, mydate.Name, "\n")
	mydate.SayHello()

}
func text(age int, name string, lastName string, n string) {
	fmt.Print(age, name, lastName, n)
	list.New().PushFront(41)
}
