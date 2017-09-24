package main

import (
	"fmt"
	"sort"
)

type people []string

func (a people) Len() int           { return len(a) }
func (a people) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a people) Less(i, j int) bool { return a[i] < a[j] }

func main() {
	studyGroup := people{"Zeno", "John", "Al", "Jenny"}
	s := []string{"Zeno", "John", "Al", "Jenny"}
	n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}
	sort.Sort(studyGroup)
	sort.StringSlice(s).Sort()
	sort.IntSlice(n).Sort()
	fmt.Println(studyGroup, s, n)

	sort.Sort(sort.Reverse(studyGroup))
	sort.Sort(sort.Reverse(sort.StringSlice(s)))
	sort.Sort(sort.Reverse(sort.IntSlice(n)))
	fmt.Println(studyGroup, s, n)
}
