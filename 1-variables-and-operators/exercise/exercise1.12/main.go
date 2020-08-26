package main

import (
	"fmt"
	"time"
)

/*
	# Zero values
	Zero value adalah variable yang berisi default value sesuai tipe datanya.
	Berikut beberapa default value pada Go:

	bool ~> false

	numbers (integer & floats) ~> 0

	string ~> "" (empty string)

	pointers, function,
	interfaces, slices, ~> nil
	channels, and maps
*/

/*
	Kita akan mencoba mendefine variable tanpa initial value.
	Kemudian print nilai dari variable tersebut.
	Kita akan menggunakan "fmt.Printf" untuk menampilkan format value.

	Format subtitution yang paling umum:
	%v ~> Any value, digunakan ketika ingin print nilai variable tanpa memperhatikan tipe datanya
	%+v ~> Values with extra information, misalnya nilai pada struct dan field namenya
	%#v ~> Sama seperti %+v ditambah tipe datanya
	%T ~> Print tipe data variable
	%d ~> Desimal (base 10)
	%s ~> String
	%f ~> Float
*/

func main() {

	// Integer
	var count int
	fmt.Printf("Count: %#v \n", count)

	// Float
	var discount float64
	fmt.Printf("Discount: %#v \n", discount)

	// Boolean
	var debug bool
	fmt.Printf("Debug: %#v \n", debug)

	// String
	var message string
	fmt.Printf("Message: %#v \n", message)

	// Collection of string
	var emails []string // slice
	fmt.Printf("Emails: %#v \n", emails)

	// Struct
	var startTime time.Time
	fmt.Printf("Start: %#v \n", startTime)
}
