package main

import "fmt"

/*
	# A simple if statement
	If statement adalah basic dari logic pada Go. Dengan if statement kita menentukan
	apakah suatu pemrosesan dijalankan atau tidak berdasarkan boolean expression.
	Notasi if statement seperti berikut: **if <boolean expression> { <code block> }**.

	Code block akan akan dijalankan ketika boolean expression bernilai true.
	If statement hanya bisa digunakan pada function scope.
*/

func main() {
	input := 5
	if input%2 == 1 {
		fmt.Println(input, "is odd")
	}
}
