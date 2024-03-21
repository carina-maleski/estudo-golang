package main

import "fmt"

func main() {
	fmt.Print("Enter number:")
	var input uint
	fmt.Scanln(&input)

	for i := 1; i <= int(input); i++ {
		
		if i % 15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i % 3 == 0 { 
			fmt.Println("Fizz")
		} else if  i % 5 == 0 {
			fmt.Println("Buzz")
		} else { 
			fmt.Println(i)
		}
	}
}