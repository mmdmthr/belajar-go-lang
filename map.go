package main

import (
	"fmt"
)

func main() {
	// map digunakan untuk menyimpan data dengan pasangan key:value
	// ini seperti associative array di php
	// https://www.w3schools.com/go/go_maps.php

	// membuat map dengan var dan :=
	// var a = map[keyType]valueType{key1:value1, key2:value2, ...}
	// b := map[keyType]valueType{key1:value1, key2:value2, ...}

	var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
	b := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}

	fmt.Printf("a\t%v\n", a)
	fmt.Printf("b\t%v\n", b)

	// membuat map dengan make()
	// var a = make(map[KeyType]ValueType)
	var c = make(map[string]string)
	c["brand"] = "Ford"
	c["model"] = "Mustang"
	c["year"] = "1964"

	fmt.Printf("c\t%v\n", c)

    // menghapus elemen map
	delete(a, "year")
	fmt.Println(a)

	// melakukan iterasi pada map
	for k, v := range c {
       fmt.Printf("%v : %v, ", k, v)
	}
}
