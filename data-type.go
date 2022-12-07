package main

import "fmt"

func main() {
	/* 
	go memiliki 3 data tipe utama
	- bool: merepresentasikan nilai boolean (true atau false) https://www.w3schools.com/go/go_boolean_data_type.php
	- Numeric: merepresentasikan tipe integer, floating point dan complex https://www.w3schools.com/go/go_integer_data_type.php
	- string: merepresentasikan nilai string https://www.w3schools.com/go/go_string_data_type.php
	*/
	var a bool = true 		// boolean
	var a1 bool				// boolean default value false
	var b int = 5			// integer
	var c float32 = 3.14	// floating point
	var d string = "Hi!"	// string

	fmt.Println("Boolean: ", a)
	fmt.Println("Boolean: ", a1)
	fmt.Println("Integer: ", b)
	fmt.Println("Float:	  ", c)
	fmt.Println("String:  ", d)
}
