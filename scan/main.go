package main

import "fmt"

func main() {
	var name string
	fmt.Println("Please enter your name below:")
	fmt.Scan(&name)
	fmt.Printf("Hello, %v \n", name)
}
