package main

import (
	"container/list"
	"fmt"
	"time"

	"github.com/jalaali/go-jalaali"
)

func main() {
	var name string = "afshin"
	age := 45
	age = 90
	text(age, name)
	fmt.Print(jalaali.ToJalaali(time.Now().Date()))
}
func text(age int, name string) {
	fmt.Print(age, name)
	list.New().PushFront(41)
}
