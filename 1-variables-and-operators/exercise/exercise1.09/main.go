package main

import "fmt"

/*
	# Operators
	Kita menggunakan variable untuk menyimpan data aplikasi.
	Sedangkan operator digunakan untuk membangun logic aplikasi.
	Operator bisa digunakan untuk membandingkan suatu data dengan data lain misalnya: sorting produk berdasarkan nilai terbesar
	bisa juga untuk memanipulasi nilai data contohnya: menjumlahkan semua harga item pada shopping cart.

	Berikut adalah jenis-jenis operator:
	- Arithmetic operator
		Digunakan untuk operasi matematika seperti addition, substraction dan multiplication.
		Contoh: +. -, *, /
	- Comparison operator
		Digunakan untuk membandingkan dua value, hasilnya yaitu: equal, not equal, less than, or greater than.
		Contoh: ==, !=, >, <, >=, <=
	- Logical operator
		Digunakan bersama boolean value untuk melihat apakah kedua true, salah satunya true, atau kombinasi lainnya.
		Contoh: &&, ||, !
	- Address operator
		Kita akan membahas ini nanti pada section pointer.
	-	Receive operator
		Digunakan ketika bekerja menggunakan Go channel, kita akan bahas pada section berikutnya.
*/

/*
	Pada exercise ini kita akan mensimulasikan restaurant bill. Kita akan menggunakan matematika dan comparison operator.

	Kita akan menjumlahkan semua pesanan lalu menghitung tip berdasarkan presentase.
	Lalu kita akan membandingkan apakah custome akan mendapat reward atau tidak.
*/

func main() {
	// main course
	var total float64 = 2 * 13
	fmt.Println("sub: ", total)

	// drinks
	total = total + (4 * 2.25)
	fmt.Println("sub: ", total)

	// discount
	total = total - 5
	fmt.Println("sub: ", total)

	// 10% tip
	tip := total * 0.1
	fmt.Println("tip: ", tip)

	// add tip to total
	total = total + tip
	fmt.Println("total: ", total)

	// split bill
	split := total / 2
	fmt.Println("split: ", split)

	// reward every 5th visit
	visitCount := 24
	visitCount = visitCount + 1

	remainder := visitCount % 5

	if remainder == 0 {
		fmt.Println("With this visit, you've earned a reward")
	}

}
