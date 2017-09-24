package main

import "fmt"

func main() {
	fibo := fibo()
	fmt.Println(fibo)
	sum := sum_even(fibo...)
	fmt.Println(sum)
}

func fibo() []int {
	var fibo []int
	fibo = append(fibo, 1, 2)
	for i, z := 2, 0; fibo[len(fibo)-1] <= 4000000; i++ {
		z = fibo[i-1] + fibo[i-2]
		fibo = append(fibo, z)
	}
	return fibo
}

func sum_even(nums ...int) int {
	var sum int
	for _, n := range nums {
		if n%2 == 0 {
			sum += n
		}
	}
	return sum
}

// Solves https://projecteuler.net/problem=2 problem.
// I feel like i populated Fibonacci sequence in some subotimal perverted way
