package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var signal chan struct{}
var DATA = make(map[int]bool)

func first(min, max int, out chan<- int) {
	for {
		select {
		case <-signal:
			close(out)
			return
		case out <- rand.Intn(max-min) + min:
		}
	}
}

func second(in <-chan int, out chan<- int) {
	for x := range in {
		_, ok := DATA[x]
		if ok {
			signal <- struct{}{}
			close(out)
		} else {
			fmt.Print(x, " ")
			out <- x
			DATA[x] = true
		}
	}
	fmt.Println()
}

func third(in <-chan int) {
	sum := 0
	for x := range in {
		sum += x
	}
	fmt.Println(sum)
}

func main() {

	input := os.Args

	if len(input) != 3 {
		fmt.Println("Need to integer parameters.")
		os.Exit(1)
	}

	min, err1 := strconv.Atoi(input[1])
	if err1 != nil {
		fmt.Println(err1)
		os.Exit(2)
	}
	max, err2 := strconv.Atoi(input[2])
	if err2 != nil {
		fmt.Println(err1)
		os.Exit(2)
	}

	if max <= min {
		fmt.Println("The first argument must be less than the second.")
		os.Exit(3)
	}

	ch1 := make(chan int)
	ch2 := make(chan int)
	signal = make(chan struct{})

	rand.Seed(time.Now().Unix())

	go first(min, max, ch1)
	go second(ch1, ch2)
	third(ch2)
}
