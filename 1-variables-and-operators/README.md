**package main**
Ini adalah deklarasi package. Suatu project Go harus memiliki minimal satu package main. File Go yang berada dalam satu folder sebaiknya memiliki nama package yang sama.

Contoh import package dari Go standard library:
import (
  "errors"
  "fmt"
  "log"
  "math/rand"
  "strconv"
  "time
)

Go memiliki standard library yang high-quality dan comprehensive. Jika package diimport dari source lain biasanya ketika diimport menggunakan nama URL, contoh: "github.com/fatih/color.

Go memiliki module system yang mempermudah penggunaan package eksternal. Untuk menggunakannya kita harus menambahkannya ke import path. Go akan secara otomatis mendowloadnya ketika kode di build.

 var cities = []string{"Jakarta", "Surabaya", "Bandung}
 Diatas merupakan contoh deklarasi list dengan tipe slice. Ada tiga tipe list pada Go yaitu: slice, arrays, dan maps. Semuanya merupakan suatu collection dengan keys dan values, dimana key digunakan untuk mengakses value.

Fungsi adalah kode yang dijalankan ketika dipanggil. func main() merupakan entry point dari aplikasi Go.

Struct adalah collection yang berisi properties dan functions. Konsepnya mirip object pada bahasa JavaScript.

## Deklarasi multiple variable menggunakan var pada single line
Umumnya kita menggunakan shorthand untuk mendefine variable pada single line. Kita juga bisa menggunakan var instead shorthand, namun saratnya tipe datanya harus sama.

**Nama Variable Non-English**
Go merupakan UTF-8 compliant language, artinya kita bisa mendefine variable menggunakan alphabet lain selain latin, misalnya Arabic atau Chinnese. Karakter pertama pada variable harus berupa letter atau _ (underscore).

Note: Tidak semua bahasa pemrograman mengizinkan kita untuk menggunakan UTF-8 sebagai nama variable atau fungsi. Fitur ini mungkin saja menjadi salah satu alasan mengapa Go sangat popular pada negara Asian misalnya China.

---

### Summary
Pada chapter ini kita telah mempelajari banyak konsep terkait variable, seperti bagaimana variable dideclare dan notasi yang digunakan untuk mendefine variable. Kita juga mempelajari bagaimana cara mengubah nilai variable yang sudah dideclare.

Semua data di apliakasi Go disimpan di variable. Dengan data kita bisa membuat operasi logic sehingga aplikasi menjadi dinamis. Kita telah mempraktikan bagaimana membandingkan nilai dua variable menggunakan comparison.

Kita telah mengexplore bagaimana Go mengimplementasikan sistem variable seperti zero value (nilai default), pointers dan scope logic. Dengan demikian kita bisa membuat software yang efisien.

Kita juga mempelajari cara mendeclare immutable variable menggunakan constant dan bagaimana iota dapat mempermudah pembuatan list atau related constant to work, misalnya enums.

Pada chapter berikutnya kita akan mempelajari bagaimana mendefine logic dan looping pada koleksi data yang kita miliki.