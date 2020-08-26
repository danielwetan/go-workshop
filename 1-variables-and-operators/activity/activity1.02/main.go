package main

import "fmt"

// # Pointer value swap3
// Swap valu a dan b sehingga a = 5 dan b = 10

func main() {
	a, b := 5, 10
	// call swap here
	swap(&a, &b)
	fmt.Println(a == 10, b == 5)
}

func swap(a *int, b *int) {
	c := new(int) // temporary variable
	*c = *a
	*a = *b
	*b = *c
}
