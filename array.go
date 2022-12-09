package main

import (
	"fmt"
)

func main() {
	/* https://www.w3schools.com/go/go_arrays.php
		array digunakan untuk menyimpan beberapa nilai 
		yang sama dalam satu variabel, untuk mendeklarasikannya
		var array_name = [length]datatype{values} jumlah item didefinisikan
		var array_name = [...]datatype{values} jumlah item tidak didefinisikan
	*/
	var arr1 = [3]int{1,2,3}
	arr2 := [5]int{4,5,6,7,8}
	var cars = [...]string{"Volvo", "BMW", "Ford", "Mazda"}

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(cars)

	/* mengakses array dapat dilakukan dengan nomor index */
	fmt.Println(cars[1])

	/* mengubah nilai elemen array */
	cars[2] = "Toyota"
	fmt.Println(cars)

	/* mencari jumlah elemen array */
	fmt.Println(len(cars))
	
}
