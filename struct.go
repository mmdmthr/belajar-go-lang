package main

import "fmt"

// struct (structure) digunakan
// untuk membuat koleksi dari beberapa tipe data
// ke dalam satu variabel

// mendeklarasikan struct
    type Person struct {
        name string
        age int
        job string
        salary int
    }

func main() {
    var pers Person
    pers.name = "David"
    pers.age = 45
    pers.job = "Teacher"
    pers.salary = 7000

    // Access and print Pers info
    fmt.Println("Name: ", pers.name)
    fmt.Println("Age: ", pers.age)
    fmt.Println("Job: ", pers.job)
    fmt.Println("Salary: ", pers.salary)
}
