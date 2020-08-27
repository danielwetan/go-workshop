package main

import (
	"fmt"
	"time"
)

// # Experssionless switch statements

func main() {
	switch dayBorn := time.Saturday; {
	case dayBorn == time.Saturday || dayBorn == time.Sunday:
		fmt.Println("Born on the weekend")
	default:
		fmt.Println("Born some other day")
	}
}
