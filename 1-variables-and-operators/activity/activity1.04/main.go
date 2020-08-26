package main

import "fmt"

/*
	# Bad count bug

	Code berikut harus menghasilkan true, namun hasil dibawah false.
	Mari kita fix bug pada code dibawah.

	Scope variable pada Go:
	- package
	- function

	func main() {
	count := 0
	if count < 5 {
		count := 10
		count++
	}
	fmt.Println(count == 11)

*/

func main() {
	count := 0
	if count < 5 {
		count = 10
		count++
	}
	fmt.Println(count == 11)
}
