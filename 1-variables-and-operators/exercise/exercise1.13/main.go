package main

import (
	"fmt"
	"time"
)

func main() {
	// membuat pointer menggunakan var
	var count1 *int // value default akan di set nil

	// membuat pointer menggunakan new
	count2 := new(int) // memiliki value default integer

	// membuat pointer dari variable
	countTemp := 5
	count3 := &countTemp
	t := &time.Time{}

	fmt.Printf("count1: %#v \n", count1)
	fmt.Printf("count2: %#v \n", count2)
	fmt.Printf("count3: %#v \n", count3)
	fmt.Printf("time: %#v \n", t)
}
