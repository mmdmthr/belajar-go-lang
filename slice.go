package main

import (
	"fmt"
)

func main() {
	/* slice mirip dengan array, namun lebih fleksibel
	tidak seperti array yang panjangnya tetap, nilai elemen slice 
	dapat ditambah dan dikurangi 

	ada 3 cara membuat slice
	- dengan format []datatype{value}
	- buat slice dari array
	- buat slice dengan fungsi make()
	*/

	// dengan format []datatype{value}
	myslice := []int{} // mendeklarasikan slice dengan panjang 0 dan kapasitas 0
	// panjang adalah berapa nilai yang sedang ditampung len()
	// kapasitas adalah jumlah maksimum variable yang dapat ditampung cap()

	fmt.Println(len(myslice))
	fmt.Println(cap(myslice))

	// buat slice dari array
	// var myarray = [length]datatype{value}
	// myslice := myarray[start:end]
	arr1 := [6]int{10, 11, 12, 13, 14, 15}
	myslice1 := arr1[2:4]

	fmt.Printf("myslice = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))

	// buat slice dengan fungsi make()
	// slice_name := make([]type, length, capacity)
	myslice2 := make([]int, 5, 10)
	fmt.Printf("myslice1 = %v\n", myslice2)
	fmt.Printf("length = %d\n", len(myslice2))
	fmt.Printf("capacity = %d\n", cap(myslice2))

	// memasukkan nilai ke dalam slice dapat dilakukan dengan fungsi append()
}
