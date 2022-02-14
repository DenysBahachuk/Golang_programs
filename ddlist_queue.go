package main

import "fmt"

type Node struct {
	Prev  *Node
	Value int
	Next  *Node
}

var Queue = new(Node)

func traverse(root *Node) {
	if root == nil {
		fmt.Println("The Queue is empty")
		return
	}
	for root != nil {
		fmt.Print(root.Value, " -> ")
		root = root.Next
	}
	fmt.Println()
}

func Push(root *Node, v int) {
	if root == nil {
		root := &Node{nil, v, nil}
		Queue = root
		return
	}
	newNode := &Node{nil, v, root}
	Queue = newNode
}

func Pop(root *Node) int {
	if root == nil {
		fmt.Println("The Queue is empty")
		return -1
	}
	if root.Next == nil {
		value := root.Value
		root = nil
		return value
	}
	for root.Next.Next != nil {
		root = root.Next
	}
	value := root.Next.Value
	root.Next = nil
	return value
}

func main() {

	Push(Queue, 10)
	Push(Queue, -3)
	Push(Queue, -2)
	traverse(Queue)
	Pop(Queue)
	Pop(Queue)
	Pop(Queue)
	traverse(Queue)
}
