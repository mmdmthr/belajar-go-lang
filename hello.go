package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("hello, world\n")
	fmt.Println(os.Getenv("USER"), ", Let's be friends!")
}
