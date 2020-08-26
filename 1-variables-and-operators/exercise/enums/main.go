package main

/*
# Enums

Enums adalah cara mendeklarasikan fixed list yang memilih value yang saling terkait.
Go tidak mempunyai built-in tipe untuk enums, namun kita bisa membuatnya dengan iota
untuk mendefine enums menggunakan constants.

Contohnya code dibawah merupakan nama hari dalam sepekan yang didefine menggunakan constants.
Code seperti bisa diterapkan fitur "iota" Go.

const (
	Sunday    = 0
	Monday    = 1
	Tuesday   = 2
	Wednesday = 3
	Thursday  = 4
	Friday    = 5
	Saturday  = 6
)

Berikut adalah code yang sama, didefine menggunakan iota
data dibawah menggunakan iota number
dengan iota, enums lebih mudah dibuat dan dimantain, terutama ketika menambahkan data baru ditengah code nantinya.
*/

const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)
