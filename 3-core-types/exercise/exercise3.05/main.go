package main

import "fmt"

// # Safely looping over a string

func main() {
	logLevel := "デバッグ"

	for index, runeVal := range logLevel {
		fmt.Println(index, string(runeVal), runeVal)
	}
}
