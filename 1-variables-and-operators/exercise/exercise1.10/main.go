package main

/*
	Shorthand operator
	-- ~> reduce value by 1
	++ ~> increase value by 1
	+= ~> add and assign
	-= ~> subtract and assign
*/

import "fmt"

func main() {
	count := 5
	count += 5
	fmt.Println(count)

	count++
	fmt.Println(count)

	count--
	fmt.Println(count)

	count -= 5
	fmt.Println(count)

	count += 5
	fmt.Println(count)

	name := "Daniel"
	name += " Saputra"
	fmt.Println("Hello", name)
}
