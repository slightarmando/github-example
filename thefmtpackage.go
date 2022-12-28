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
	fmt.Println("Prints mixed value")
	fmt.Printf("%b\t%x\t%#x\n", y, y, y)

	foo()
}

func foo() {
	s := fmt.Sprintf("%b\t%x\t%#x\n", y, y, y)
	fmt.Println("This uses Sprintf and ASSIGNS Var S with string of chars below")
	fmt.Println(s)
}
