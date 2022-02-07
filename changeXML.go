package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Record struct {
	Numbers []int `xml:"Numbers>value"`
}

func main() {
	//provide an XML file as a command line argument
	input := os.Args

	if len(input) != 2 {
		fmt.Println("Provide one XML file!")
		return
	}

	fileName := input[1]

	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading file. ", err)
		return
	}

	data := Record{}

	err1 := xml.Unmarshal(file, &data)
	if err1 != nil {
		fmt.Println("Error unmarshaling file. ", err1)
		return
	}

	for i, v := range data.Numbers {
		data.Numbers[i] = v + 1
	}

	fileXML, err2 := xml.MarshalIndent(data, "", "     ")
	if err2 != nil {
		fmt.Println(err2)
		return
	}

	err3 := ioutil.WriteFile(fileName, fileXML, 0)
	if err3 != nil {
		fmt.Println(err3)
		return
	}
}
