package main

import "fmt"

func main() {
	for x := 1; x < 10; x++ {
		fmt.Println(x)
	}

	foo()
}

func foo() {

	y := 1
	fmt.Println((y))

}
