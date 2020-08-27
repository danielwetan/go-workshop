package main

import "fmt"

// # Looping over arrays and slices

func main() {
	names := []string{"Jim", "Jane", "Joe", "June"}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}
}
