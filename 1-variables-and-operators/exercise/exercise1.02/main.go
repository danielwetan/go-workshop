package main

import "fmt"

// package level scope
// dapat diakses oleh semua fungsi di file ini
var name string = "Daniel"

func main() {
	// # deklarasi variable menggunakan var
	/*
		var ~> deklarasi variable
		foo ~> nama variable
		string ~> tipe data
		"Daniel" ~> initial value
	*/

	// function level scope
	// hanya dapat diakses di fungsi main
	var city string = "Jakarta"

	fmt.Println(name, city)
}
