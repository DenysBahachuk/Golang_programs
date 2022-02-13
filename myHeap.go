package main

import (
	"container/heap"
	"fmt"
)

type Items []interface{}

func (i *Items) Push(x interface{}) {
	item := x.(interface{})
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
	switch i[x].(type) {
	case int:
		return i[x].(int) < i[y].(int)
	case float64:
		return i[x].(float64) < i[y].(float64)
	case string:
		return i[x].(string) < i[y].(string)
	default:
		fmt.Println("Type not defined")
	}
	return false
}

func (i Items) Swap(x, y int) {
	i[x], i[y] = i[y], i[x]
}

func main() {

	items := Items{2, 4, 5, -3, 0}

	fmt.Print(items, " ")
	fmt.Println()
	heap.Init(&items)
	fmt.Print(items, " ")
}
