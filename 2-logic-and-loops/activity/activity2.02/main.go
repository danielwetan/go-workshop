package main

import "fmt"

/*
	- Write a program that prints out the numbers from 1 to 100
	- If the number is a multiple of 3, print "Fizz"
	- If the number is a multiple of 5, print "Buzz"
	- If the number is a multiple of 3 and 5, print "FizzBuzz"
*/

func main() {
	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
