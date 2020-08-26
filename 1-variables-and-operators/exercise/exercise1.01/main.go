package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // seed random number generator
	r := rand.Intn(5) + 1            // generate angka 0 - 4, lalu tambahkan 1 agar nilanya menjadi 1-5
	starts := strings.Repeat("*", r) // gunakan string repeater sebanyak r kali.
	fmt.Println(starts)              // print starts ke console
}
