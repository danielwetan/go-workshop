package main

import (
	"fmt"
	"math"
	"math/big"
)

// # Big numbers

func main() {
	intA := math.MaxInt64
	intA = intA + 2

	bigA := big.NewInt(math.MaxInt64)
	bigA.Add(bigA, big.NewInt(2))

	fmt.Println("MaxInt64: ", math.MaxInt64)
	fmt.Println("Int:", intA)
	fmt.Println("Big Int:", bigA.String())

}
