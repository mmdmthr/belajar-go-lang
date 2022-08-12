package main

import (
	"fmt"
)

func main() {
	// fungsi Print() mencetak argument dengan format default
	var i, j string = "Hello", "World"
	fmt.Print(i)
	fmt.Print(j)
	fmt.Print("\n") // mencetak baris baru

	// dapat mencetak beberapa argumen yang diinput
	fmt.Print(i, " ", j)
	fmt.Print("\n")

	// otomatis menambahkan spasi jika tidak ada variabel string
	var d, f = 10, 20
	fmt.Print(d, f)
	fmt.Print("\n")

	// fungsi Println() hampir sama dengan Print() dengan perbedaan
	// menambahkan spasi di antara argumen dan menambah baris baru di akhir baris
	fmt.Println(i, j)
}
