package main

import "fmt"

func main() {
	fmt.Print("Enter number:")
	var input uint
	fmt.Scanln(&input)

	for i := 1; i <= int(input); i++ {
		
		switch {
		case i%15 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}

	}
}