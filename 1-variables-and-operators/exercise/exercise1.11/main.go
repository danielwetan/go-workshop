package main

import "fmt"

/*
	 # Comparing values
	 Comparison operator:
		== ~> TRUE jika nilai kedua value sama
		!= ~> TRUE jika nilai kedua value berbeda
		< ~> TRUE jika value dikiri lebih kecil daripada kanan
		<= ~> TRUE jika value dikiri lebih kecil atau sama dengan value dikanan
		> ~> TRUE jika value dikiri lebih besar daripada kanan
		>= ~> TRUE jika value dikiri lebih besar atau sama dengan value dikanan

		Logical operator:
		&& ~> TRUE jika kedua value true
		|| ~> TRUE jika salah satu value true
		! ~> kebalikan / negasi, mengubah true menjadi false, false menjadi true
*/

func main() {
	/*
		Kita akan membuat program untuk mengecek level membership pengunjung
		berdasarkan jumlah visit ke restaurant.

		Level membership:
		- Silver: antara 10 dan 20 visits
		- Gold: antara 21 dan 30 visits
		- Platinum: lebih dari 30 visits
	*/

	visits := 15
	fmt.Println("First visit:", visits == 1)
	fmt.Println("Return visit:", visits != 1)
	fmt.Println("Silver member:", visits >= 10 && visits < 21)
	fmt.Println("Gold member:", visits > 20 && visits <= 30)
	fmt.Println("Platinum member:", visits > 30)
}
