package main

import "fmt"

func main() {
	
	/* integers, floats */
	fmt.Println("1 + 1 =", 1+1) // integers
	fmt.Println("1.0 + 1.0 =", 1.0+1.0) // floating numbers

	/* strings */
	fmt.Println(len("Hello World")) // string length like python :)
	fmt.Println("Hello World"[1]) // 101
	fmt.Println("Hello " + "World") // string concenate

	/* boolean */
	fmt.Println("True && True =", true && true)
	fmt.Println("True && False =", true && false)
	fmt.Println("True || True =", true || true)
	fmt.Println("True || False =", true || false)
	fmt.Println("!true =", !true)
}
