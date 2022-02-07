package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	//provide a JSON file as a command line argument
	input := os.Args

	if len(input) != 2 {
		fmt.Println("Provide one JSON file!")
		return
	}

	fileName := input[1]

	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading file. ", err)
		return
	}

	var data []interface{}
	json.Unmarshal(file, &data)

	for i, d := range data {
		t := fmt.Sprint(d)
		n, err := strconv.Atoi(t)

		if err != nil {
			fmt.Println(err)
			return
		}
		data[i] = n + 1
	}

	fileJSON, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	err1 := ioutil.WriteFile(fileName, fileJSON, 0)
	if err != nil {
		fmt.Println(err1)
		return
	}
}
