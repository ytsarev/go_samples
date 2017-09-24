package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(greatest(3, 6, 55, 15, 23, 8))
}

func greatest(nums ...int) int {
	sort.Ints(nums)
	return nums[len(nums)-1]
}
