package main

import (
	"fmt"
)

// TIPE VARIABEL PADA GO

// int- menyimpan integer (angka), contoh: 123 atau -123
// float32- menyimpan angka dengan koma (desimal): contoh: 19.99 atau -19.99
// string - menyimpan teks, contoh: "Hello World". Nilai string ditulis dalam dua tanda petik ("")
// bool- menyimpan nilai kondisi: true atau false

// MENDEKLARASIKAN VARIABE

// menggunakan kata kunci var (harus menentukan tipe, nilai, atau keduanya)
// var variablename type
// di Go semua variabel diinisialisasi. Jika mendeklarasikan variabel tanpa nilai awal,
// nilainya akan disetel ke nilai default tipenya:
var a string // nilai otomatis "" (string kosong)
var b int    // nilai otomatis 0
var c bool   // nilai otomatis false

// var variablename value
var j = 10

// var variablename type = value
var k int = 10
var msg string = "Remote host found"

func main() {
	// Penetapan Nilai Setelah Deklarasi, dilakukan di dalam fungsi
	b = 2

	// menggunakan tanda :=
	// tipe data diturunkan dari nilai variabelnya, compiler akan menentukan secara otomatis
	// hanya dapat dilakukan di dalam fungsi
	// harus menetapkan nilai variabel

	// variablename := value
	name := "John"
	age := 25

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(name)
	fmt.Println(age)
}

// Perbedaan antara var dengan :=
// Var dapat digunakan di dalam dan di luar fungsi sementara := hanya dapat digunakan di dalam fungsi
// Var dapat melakukan deklarasi variabel dan penetapan nilai secara terpisah
// sementara dengan := tidak dapat melakukan deklarasi variabel secara terpisah (harus dilakukan pada baris yang sama)
