package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func delSpaces(filePath string) error {
	file, err := os.OpenFile(filePath, os.O_RDWR, 0)
	if err != nil {
		fmt.Println("Error opening file")
		return err
	}

	reader := bufio.NewReader(file)
	var DATA string
	for {
		str, err1 := reader.ReadString('\n')
		if err1 == io.EOF {
			break
		} else if err1 != nil {
			return err1
		}
		str = strings.ReplaceAll(str, " ", "")
		DATA = DATA + str
	}
	file.Close()

	err2 := os.Remove(filePath)
	if err2 != nil {
		fmt.Println("Error removing the file")
		return err2
	}

	newFile, err3 := os.Create(filePath)
	defer newFile.Close()
	if err3 != nil {
		fmt.Println("Error creating a new file")
		return err3
	}

	_, err4 := newFile.WriteString(DATA)
	if err4 != nil {
		fmt.Println("Error writing to a new file")
		return err4
	}
	return nil
}

func main() {

	input := os.Args
	if len(input) < 2 {
		fmt.Println("Input a path to a file!")
		os.Exit(1)
	}
	filePath := input[1]

	err := delSpaces(filePath)
	if err != nil {
		fmt.Println(err)
	}
}
