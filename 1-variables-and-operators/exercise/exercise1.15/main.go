package main

import "fmt"

/*
	# Function design with pointers
	Ketika value dengan pointer digunakan sebagai parameter function, Go tidak akan membuat saliannnya.
	Sehingga perubahan yang dilakukan pada value tersebut akan berdampak pada value diluar function.
*/

func addValue(count int) {
	count += 5
	fmt.Println("addValue:", count)
}

func addPoint(count *int) {
	*count += 3
	fmt.Println("addPoint:", *count)
}

func main() {
	var count int
	addValue(count) // value asli tidak berubah
	fmt.Println("addValue main:", count)

	addPoint(&count) // value asli berubah
	fmt.Println("addPoint main:", count)
}
