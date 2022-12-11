package main

import (
	"fmt"
)

// fungsi adalah sebuah blok kode
// yang tidak dieksekusi langsung
// fungsi adalah blok kode yang disimpan
// untuk dapat dijalankan saat dipanggil
	func myMessage() {
		fmt.Println("I just got executed!")
	}

	func sum(x int, y int) (result int) {
		result = x + y
		return
	}

	func helloWorld(x int, y string) (result1 int, txt1 string) {
		result1 = x + x
		txt1 = y + " World!"
		return
	}
	
func main() {
	myMessage()

	// parameter 
	// namaFungsi(param1 type)

	// fungsi dalam go dapat mengembalikan nilai dengan typenya
	fmt.Println(sum(1,2))

	// fungsi juga dapat mengembalikan lebih dari satu nilai
	fmt.Println(helloWorld(5, "Hello"))

	// cara untuk menyimpan dalam variable
	a, b := helloWorld(5, "Hello")
	fmt.Println(a, b)
}
