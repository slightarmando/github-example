package main

import "fmt"

var a int

type hotdog float32

var b hotdog

var y int

func main() {
	a = 42
	b = 19.0
	fmt.Printf("%T\n", a)
	fmt.Println(a)
	fmt.Printf("%T\n", b)
	fmt.Println(b)
	y = int(b)
	fmt.Printf("%T\n", y)
	fmt.Println(y)

}
