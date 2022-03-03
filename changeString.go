package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// Input path to the file and two strings as command line arguments
	file1 := flag.String("f", "", "filePath")
	str1 := flag.String("s1", "", "string1")
	str2 := flag.String("s2", "", "string2")
	flag.Parse()

	fmt.Println(flag.NFlag())
	if flag.NFlag() < 3 {
		fmt.Println("Not enough arguments!")
		os.Exit(1)
	}

	fileName := *file1
	string1 := *str1
	string2 := *str2

	file, err := os.Open(string(fileName))
	defer file.Close()
	if err != nil {
		fmt.Println("Error opening file", err)
		os.Exit(2)
	}

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		line = strings.TrimRight(line, "\r\n")
		if line == string1 {
			line = string2
		}
		fmt.Println(line)
	}
}
