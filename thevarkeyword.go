package main

import "fmt"

// Global declaration for var z and y
var y = 36
var z int

func main() {
	//short declaration
	//DECLARE a variable and ASSIGN a VALUE (of a certai type)
	x := 42
	fmt.Println(x)
	fmt.Println(y)

	foo()
}

func foo() {
	fmt.Println("Hey dude, this is the number,", y, ",well done!")
	fmt.Println("here is the other number", z)
	fmt.Println("Testing this, well done")
}
