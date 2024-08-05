package main

import "fmt"

var i int = 27 //Declares a package level integer variable

func main() {
	//Three ways to declare and initialize a integer variable
	var a int = 1
	var b int
	b = 2
	c := 3
	fmt.Println(a, b, c)

	fmt.Println(i)
	var i int = 42 //Shadowing the package level integer variable
	fmt.Println(i)

}
