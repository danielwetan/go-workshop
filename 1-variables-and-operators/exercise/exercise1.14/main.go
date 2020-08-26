package main

import (
	"fmt"
	"time"
)

/*
	# Mendapatkan value dari pointer
	Pada latihan sebelumnya kita mencoba print pointer variable ke console,
	disitu kita mendapatkan either nil atau memory address.
	Untuk mendapatkan value dari suatu pointer, gunakan * didepan nama variablenya.

	Membedakan zero dan nil pointer adalah bug umum pada aplikasi Go,
	kompiler tidak dapat warn you about it dan itu terjadi ketika aplikasi sudah berjalan.
	Best practicenya adalah selalu check pointer apakah nil atau tidak sebelum digunakan.
	Namun ada saatnya kita tidak perlu melakukan pengecekan, misalnya ketika property atau function on a struct.
	Go akan memberikan error ketika kita mengecek value yang tidak bisa di dereference.
*/

func main() {
	var count1 *int
	count2 := new(int)
	countTemp := 5
	count3 := &countTemp
	t := &time.Time{}

	// Pengecekan value pointer dengan nil
	// *count1 ~> mengambil value asli dari pointer, bukan memory addressnya
	if count1 != nil {
		fmt.Printf("count1: %#v \n", *count1)
	}
	if count2 != nil {
		fmt.Printf("count2: %#v \n", *count2)
	}
	if count3 != nil {
		fmt.Printf("count3: %#v \n", *count3)
	}
	if t != nil {
		fmt.Printf("time: %#v \n", *t)
		fmt.Printf("time: %#v \n", t.String()) // kita tidak perlu me-dereference
	}
	// With our time variable, we were able to dereference the whole struct, which is why our output doesn't start with an &.
}
