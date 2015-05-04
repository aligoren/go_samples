package main

import "fmt"

var scope string = "This is scope"

func main() {
	var x string = "Hello World"

	fmt.Println(x)

	x = "Hello World. I'm from Turkey! "

	fmt.Println(x)

	x = x + "How are you?"

	fmt.Println(x)

	var a string = "hello"
	var b string = "world"

	fmt.Println(a == b) // false

	z := "This is a string"

	fmt.Println(z)

	n :=  1 // n := x mean => as x value(string, float, int vs.)
	n1 := 2

	fmt.Println(n*n1)

	name := "Ali"

	fmt.Println("Hi! My name is", name)

	friendName := "Mehmet"

	fmt.Println("My friend's name is", friendName)

	fmt.Println(scope)

	const constant string = "This is constant"

	fmt.Println(constant)

	//constant = "Do not this! You get an error. Because this is const variable."

	/* multiple variable defination */
	var (
		number1 = 15
		number2 = 10
		text = "Multiplication is:"
	)

	fmt.Println(text,number1*number2)

	fmt.Println("Enter a number: ")
	var inp float64
	fmt.Scanf("%f", &inp)

	outp := inp * 2

	fmt.Println("Multiplication:", outp)

}

func f() {
	fmt.Println(scope) // it's work because scope is global

	//fmt.Println(x) // doesn't work because x not global. x local variable
}