/*
	# Message bug
	Kode dibawah ini tidak berjalan. Mari kita solve bug pada kode ini.

package main
import "fmt"
func main() {
	count := 5
	if count > 5 {
		message := "Greater than 5"
	} else {
		message := "Less than 5"
	}
	fmt.Println(message)
}
*/

/*
	Scope variable akan berpengaruh pada code,
	oleh karena itu rencanakan terlebih dahulu sebelum membuat code.
*/

package main

import "fmt"

func main() {
	count := 5
	message := ""
	if count > 5 {
		message = "Greater than 5"
	} else {
		message = "Less than 5"
	}
	fmt.Println(message)
}
