package main

import "fmt"

func main() {
	// int 
	// float32
	// bool
	// string

	// var a int = 10
	// var a  = 20

	// a := 10

	a := 10

	fmt.Println(a)

	if a > 18 {
		fmt.Println("you are adult")
	}else if a == 18 {
		fmt.Println("you are 18")
	}else {
		fmt.Println("you are not adult")
	}
}