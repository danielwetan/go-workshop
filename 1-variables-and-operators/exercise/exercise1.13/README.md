### Value versus Pointer
Ketika kita menggunakan value seperti int, bool dan string sebagai parameter function,
Go akan membuat copy value tersebut untuk digunakan pada function tersebut.
Dengan demikian ketika kita melakukan operasi pada value tersebut, nilai originalnya tidak akan berubah.
Melainkan value hasil copy yang berubah.

Go menggunakan simple memory management system yang disebut stack.
Sisi negatif dari mengcopy value adalah bisa memakan banyak memory, ketika kita mempassingnya ke function.
Dampaknya akan sangat terasa ketika aplikasi Go sudah kompleks dan memiliki banyak function.

Namun ada alternatif dari masalah tersebut,
Instead mempassing value secara langsung kita bisa membuat pointer sebagai pengganti value tersebut.
Pointer bukanlah value itu sendiri, melainkan penunjuk ke address dari suatu value di memory.
Dengan pointer Go tidak akan membuat salinan dari value tersebut ketika mempassingnya di function.

Ketika membuat pointer, Go tidak bisa memanage value memory menggunakan stack.
Pointer akan dihandle oleh heap. Heap akan mengizinkan value exist sampai tidak ada bagian software yang menggunakannya.

Membuat pointer artinya value tersebut akan diletakan di heap.
Proses pengecekan apakah value perlu diletakan pada heap disebut sebagai escape analysis.
Namun bisa saja terjadi kasus value tanpa pointer diletakan pada heap, dan tidak diketahui alasannya.
Kita sendiri tidak bisa mengatur apakah suatu value disimpan di stack atau heap.

Keuntungan penggunaan pointer adalah optimasi pada pengguaan memory.
Ketika suatu value dicopy, Go membutuhkan CPU cycles untuk mendapatkan memory tersebut dan mereleasenya nanti.
Dengan pointer kita bisa menghindari proses ini sehingga lebih hemat memory.

Selain untuk performance, pointer juga digunakan untuk meningkatkan desain code.
Dengan pointer kita bisa membuat interfaces yang clean dan code yang simple.
Contoh: Kita ingin mengecek apakah value memiliki nilai atau tidak, value tanpa pointer
akan memiliki nilai default, sehingga ketika dicek hasilnya akan valid meskipun kita belum menginisalisasinya.
Dengan pointer kita bisa mendapatkan hasil yang lebih akurat karena value akan diset menjadi nil.
Pada Go, nil adalah special type yang merepresentasikan sesuatu yang tidak memiliki nilai.

Ketika kita mencoba mengambil value dari pointer kosong, hasilnya adalah runtime error.
Oleh karena itu, biasanya kita membandingkan pointer dengan nil sebelum mengambil valuenya.
Contoh: <pointer> != nil. Kita juga bisa membandingkan pointer dengan pointer lainnya, namun hasil true akan didapatkan ketika membandingkan dengan valuenya sendiri.
