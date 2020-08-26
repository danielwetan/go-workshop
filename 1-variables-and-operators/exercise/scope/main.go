package main

import "fmt"

/*
	# Scope

	Variable pada Go memiliki scope. Top level scope adalah package.
	Suatu scope bisa mempunyai child scope didalamnya.

	Ada beberapa cara mendefine child scope, cara paling mudah adalah
	mendefine scope baru didalam { }, misalnya menggunakan function, if statement dll.

	Ketika mengakses variable, Go akan mencari pada scope tempat didefine.
	Jika tidak ada maka akan dicari pada parent scope, kemudian grandparent scope dan seterusnya hingga package scope.
	Go akan terus mencari sampai ke scope package, jika ditemukan maka akan direturn pesan error.
*/

var level = "pkg"

func main() {
	fmt.Println("Main start:", level)
	// Create shadow variable
	level := 42
	if true {
		fmt.Println("Block start:", level)
		funcA()
	}
	fmt.Println("Main end:", level)
}

func funcA() {
	fmt.Println("funcA start:", level)
}
