package main

import "fmt"

func changeSlice(s *[]int) {

	for i := range *s {
		(*s)[i] = (*s)[i] + 1
	}
	*s = append(*s, 12)
}

func main() {

	sl := []int{1, 3, 5, 8, 10}
	fmt.Println(sl)
	changeSlice(&sl)
	fmt.Println(sl)
}
