package main

import "fmt"

func main() {
	
	/* we want this 1 to 10 :D
	fmt.Println(`1
		2
		3
		4
		5
		6
		7
		8
		9
		10
		`)*/

	/*i := 1

	for i <= 10 {
		fmt.Println(i) // try without i = i + 1 statemens :D
		i = i + 1
	}*/

	for i := 1; i <= 10; i++ {
		// fmt.Println(i) // It's better from above.

		/*if i % 2 == 0 {
			fmt.Println(i, "odd")
		} else {
			fmt.Println(i, "even")
		}*/

		/* i thinking this is bad practice
		if i == 0 {
			fmt.Println("Zero")
		} else if i == 1 {
			fmt.Println("One")
		} else if i == 2 {
			fmt.Println("Two")
		} else if i == 3 {
			fmt.Println("Three")
		} else if i == 4 {
			fmt.Println("Four")
		} else if i == 5 {
			fmt.Println("Five")
		} else if i == 6 {
			fmt.Println("Six")
		} else if i == 7 {
			fmt.Println("Seven")
		} else if i == 8 {
			fmt.Println("Eight")
		} else if i == 9 {
			fmt.Println("Nine")
		} else if i == 10 {
			fmt.Println("Ten")
		}*/

		switch i { /* good practice but only for study*/
			case 0: fmt.Println("Zero")
			case 1: fmt.Println("One")
			case 2: fmt.Println("Two")
			case 3: fmt.Println("Three")
			case 4: fmt.Println("Four")
			case 5: fmt.Println("Five")
			case 6: fmt.Println("Six")
			case 7: fmt.Println("Seven")
			case 8: fmt.Println("Eight")
			case 9: fmt.Println("Nine")
			case 10: fmt.Println("Ten")
			default: fmt.Println("Unknown number")
		}
	}
}
