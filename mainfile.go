package main

import "fmt"

func main() {
	name := "Jeff"
	fmt.Printf("Welcome %v, how are you doing?\n", name)

	var age int

	fmt.Scan(&age)
	fmt.Printf("Wow, you are %v years old!", age)
}
