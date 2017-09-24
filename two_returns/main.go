package main

import "fmt"

func main() {
	a, b := half_even(2)
	c, d := half_even(5)
	fmt.Println(a, b, c, d)
}

func half_even(num int) (int, bool) {
	return num / 2, num%2 == 0
}
