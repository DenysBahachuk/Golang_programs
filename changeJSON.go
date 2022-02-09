package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//provide a JSON file as a command line argument
	input := os.Args

	if len(input) != 2 {
		fmt.Println("Provide one JSON file!")
		os.Exit(1)
	}

	fileName := input[1]

	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading file. ", err)
		os.Exit(2)
	}

	var target Target

	err1 := json.Unmarshal(file, &target)
	if err1 != nil {
		fmt.Println(err)
		os.Exit(3)
	}

	for i, v := range target.Numbers {
		target.Numbers[i].Int = v.Int + 1
	}

	fileJSON, err := json.Marshal(target)
	if err != nil {
		fmt.Println(err)
		os.Exit(4)
	}

	err2 := ioutil.WriteFile(fileName, fileJSON, 0)
	if err2 != nil {
		fmt.Println(err1)
		os.Exit(5)
	}
}
