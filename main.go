package main

import "fmt"

func main() {

	for i := 0; i < 3; i++ {
		text("s")
		fmt.Print("hhhh")
	}

	// text("What")
	// fmt.Print("hhhh")
}
func text(t string) {
	fmt.Print(t, "444")
}
