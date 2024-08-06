package main

import "fmt"

var i int = 27 //Declares a package level variable

func main() {
	fmt.Println(i)
	var i int = 42 //Shadowing package level variable (Function level variable)
	fmt.Println(i)

	//Three ways to declare and initialize a integer variable
	var a int = 1
	var b int
	b = 2
	c := 3
	fmt.Println(a, b, c)

	//Float 64 vs Float 32
	var floatNum1 float32 = 123456.789 //Prints 123456.79
	var floatNum2 float64 = 123456.789 //Prints 123456.789
	fmt.Println(floatNum1, floatNum2)

	//Operations with 2 variables of different types
	var d int = 12
	var e float32 = 12.34
	//Addition
	fmt.Println(d + int(e))     //Prints 24
	fmt.Println(float32(d) + e) //Prints 24.34

	//Subtraction
	fmt.Println(d - int(e))     //Prints 0
	fmt.Println(float32(d) - e) //Prints -0.34

	//Multiplication
	fmt.Println(d * int(e))     //Prints 144
	fmt.Println(float32(d) * e) //Prints 148.08

	//Division
	fmt.Println(d / int(e))     //Prints 1
	fmt.Println(float32(d) / e) //Prints 0.97

	//Modulus
	fmt.Println(d % int(e)) //Prints 0
	//Modulus does not work with floats fmt.Println(float32(d) % e)

}
