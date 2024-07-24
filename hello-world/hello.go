package main
import "fmt"

func main(){

	fmt.Println("Hello World") //Prints 'Hello World'
	fmt.Println("Hello" + "World") //Prints 'HelloWorld'

	fmt.Println(Hello())
	
//Two ways to declare and initialize a variable
//#1
	var helloWorld string
	helloWorld = "Hello World!"
	fmt.Println(helloWorld)
//#2
	helloworld := "Hello World!" 
	fmt.Println(helloworld)
}

func Hello() string { //Declares a new function, returning a string 
	return "Hi World!" 
}

func Hello2() {
	x := 2
	y := 1 
	if x > y {
		fmt.Println("Hello World")
	} else {
		fmt.Println("Bye World")
	}
}