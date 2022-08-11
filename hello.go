package main

// Di Go, setiap program adalah bagian dari sebuah package.
// Kita mendefinisikan ini menggunakan kata kunci package.
// Dalam contoh ini, program termasuk dalam paket main.

import (
	"fmt"
)

// Import ("fmt") memungkinkan kita mengimpor file yang disertakan dalam paket fmt.
// Go mengabaikan white space seperti enter.

func main() {
	// func main() {} adalah sebuah fungsi. Kode di dalam tanda kurung kurawalnya {} akan dieksekusi.
	fmt.Printf("Hello, World!")
	// fmt.Println() adalah fungsi yang tersedia dari paket fmt.
	// Digunakan untuk menampilkan/mencetak teks. Dalam contoh kita ini akan menampilkan "Hello World!".
}
