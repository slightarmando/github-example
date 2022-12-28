package main

import "fmt"

var a int

type hotdog int

var b hotdog

func main() {
	fmt.Println("GO is all about TYPE")
	fmt.Println("You can create your own Type in GO")
	a = 42
	b = 43
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
}
