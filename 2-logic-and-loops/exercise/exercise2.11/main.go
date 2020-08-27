package main

import (
	"fmt"
	"math/rand"
)

// # Using break and continue to control loops

func main() {
	for {
		r := rand.Intn(8)

		if r%3 == 0 {
			fmt.Println("Skip")
			continue
		} else if r%2 == 0 {
			fmt.Println("Stop")
			break
		}

		fmt.Println(r)

	}
}
