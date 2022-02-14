package main

import (
	"container/heap"
	"fmt"
)

type Item struct {
	value int
}

type Items []*Item

func (i *Items) Push(x interface{}) {
	item := x.(*Item)
	*i = append(*i, item)
}

func (i *Items) Pop() interface{} {
	old := *i
	len := len(old)
	x := old[len-1]
	old[len-1] = nil
	*i = old[0 : len-1]
	return x
}

func (i Items) Len() int {
	return len(i)
}

func (i Items) Less(x, y int) bool {
	return i[x].value < i[y].value
}

func (i Items) Swap(x, y int) {
	i[x], i[y] = i[y], i[x]
}

func main() {

	items := Items{&Item{1}, &Item{-2}, &Item{3}}

	for _, v := range items {
		fmt.Print(*v, " ")
	}

	fmt.Println()
	heap.Init(&items)

	for _, v := range items {
		fmt.Print(*v, " ")
	}
	fmt.Println()

	heap.Push(&items, &Item{-5})
	heap.Push(&items, &Item{2})

	for _, v := range items {
		fmt.Print(*v, " ")
	}
	fmt.Println()

	heap.Pop(&items)

	for _, v := range items {
		fmt.Print(*v, " ")
	}
	fmt.Println()
}
