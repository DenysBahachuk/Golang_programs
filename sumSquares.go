package main

import (
	"fmt"
	"os"
	"strconv"
)

func getNumbers(min, max int, out chan<- int) {
	for i := min; i <= max; i++ {
		out <- i
	}
	close(out)
}

func getSquares(in <-chan int, out chan<- int) {
	fmt.Print("Squares: ")
	for x := range in {
		fmt.Print(x*x, " ")
		out <- x * x
	}
	fmt.Println()
	close(out)
}

func sumSquares(in <-chan int) {
	sum := 0
	for x := range in {
		sum += x
	}
	fmt.Print("Sum of the squares: ", sum)
}

func main() {
	//Input min and max integer variables of the range. Input data must be command line arguments.
	input := os.Args
	if len(input) < 3 {
		fmt.Println("Not enough arguments!")
		os.Exit(1)
	}

	min, err1 := strconv.Atoi(input[1])
	if err1 != nil {
		fmt.Println(err1)
		os.Exit(2)
	}
	max, err2 := strconv.Atoi(input[2])
	if err2 != nil {
		fmt.Println(err2)
		os.Exit(2)
	}

	if min >= max {
		fmt.Println("First argument must be less then the second!")
		os.Exit(3)
	}

	ch1 := make(chan int)
	ch2 := make(chan int)

	go getNumbers(min, max, ch1)
	go getSquares(ch1, ch2)
	sumSquares(ch2)
}
