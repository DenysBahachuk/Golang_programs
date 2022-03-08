package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"sync"
)

func readFiles(wg *sync.WaitGroup, out chan<- string, files []string) {
	defer wg.Done()
	for _, file := range files {
		f, err := ioutil.ReadFile(file)
		if err != nil {
			panic(err)
		}
		sFile := strings.ReplaceAll(string(f), "\r\n", " ")
		out <- sFile
	}
	close(out)
}

func matchFile(wg *sync.WaitGroup, in <-chan string, out chan<- int, phrase string) {
	defer wg.Done()
	countFiles := 0
	for x := range in {
		matches := strings.Count(x, phrase)
		countFiles++
		fmt.Printf("Number of matches in the %d file - %d.\n", countFiles, matches)
		out <- matches
	}
	close(out)
}

func countMatches(wg *sync.WaitGroup, in <-chan int) {
	defer wg.Done()
	sum := 0
	for x := range in {
		sum += x
	}
	fmt.Printf("Total number of matches - %d.\n", sum)
}

func main() {

	input := os.Args
	if len(input) < 3 {
		fmt.Println("Not enough arguments! Enter paths to files first and then a string to find matches.")
		os.Exit(1)
	}
	phrase := input[len(input)-1]
	files := input[1 : len(input)-1]

	ch1 := make(chan string)
	ch2 := make(chan int)

	var wg sync.WaitGroup
	wg.Add(3)

	go readFiles(&wg, ch1, files)
	go matchFile(&wg, ch1, ch2, phrase)
	go countMatches(&wg, ch2)

	wg.Wait()
}
