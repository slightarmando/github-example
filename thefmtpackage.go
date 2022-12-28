package main

import "fmt"

var y = 42

func main() {
	fmt.Println("Prints value")
	fmt.Println(y)
	fmt.Println("Prints TYPE")
	fmt.Printf("%T\n", y)
	fmt.Println("Prints Binary value")
	fmt.Printf("%b\n", y)
	fmt.Println("Prints hexidecimal value")
	fmt.Printf("%x\n", y)
	fmt.Println("Prints hexidecimal with zero x before it")
	fmt.Printf("%#x\n", y)
}
