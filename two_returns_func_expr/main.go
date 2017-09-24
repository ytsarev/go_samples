package main

import "fmt"

func main() {
	half_even := func(num int) (int, bool) {
		return num / 2, num%2 == 0
	}
	a, b := half_even(2)
	c, d := half_even(5)
	fmt.Println(a, b, c, d)
}
