// Comment adalah teks yang akan diabaikan ketika eksekusi kode
// Comment dapat digunakan untuk menjelaskan kode dan membuatnya lebih mudah dibaca
// Comment juga dapat digunakan untuk mencegah kode dieksekusi ketika melakukan tes pada kode alternatif
// Go mendukung comment dalam satu baris (single-line) dan multi baris (multi-line)

// Comment satu baris dimulai dengan dua tanda garis miring (//)
// semua teks setelah // sampai akhir baris akan diabaikan oleh kompiler (tidak akan dieksekusi)

// Ini adalah comment
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World") // Ini adalah comment

	// Comment multi baris diawali dengan /* dan diakhiri dengan */
	// Semua tekse di antara /* dan */ akan diabaikan oleh kompiler

	/* kode di bawah ini akan mencetak teks Hello worl
	ke layar */
	fmt.Println("Hello World!")

	// contoh comment untuk mencegah kode dieksekusi
	// fmt.Println("Baris ini tidak dieksekusi")
}
