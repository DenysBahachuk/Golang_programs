package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"os"
	"strings"
)

type Element struct {
	Name    string
	Surname string
	Id      string
}

var DATA = make(map[string]Element)

func print() {
	if len(DATA) == 0 {
		fmt.Println("The storage is empty")
	}
	for k, v := range DATA {
		fmt.Println(k, v)
	}
}

func lookup(key string) *Element {
	v, ok := DATA[key]
	if !ok {
		return nil
	}
	return &v
}

func add(key string, e Element) bool {
	if key == "" {
		fmt.Println("The must not be empty")
		return false
	}
	if lookup(key) != nil {
		fmt.Println("The key already exists.")
		return false
	}
	DATA[key] = e
	return true
}

func change(key string, e Element) bool {
	if lookup(key) == nil {
		return false
	}
	DATA[key] = e
	return true
}

func del(key string) bool {
	if lookup(key) == nil {
		return false
	}
	delete(DATA, key)
	return true
}

func load(path string) error {
	fmt.Println("Loading:", path)

	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		fmt.Println("The file doesn't exist.")
		return err
	}

	decoder := gob.NewDecoder(file)
	err1 := decoder.Decode(&DATA)
	if err1 != nil {
		fmt.Println("Error decoding the file")
		return err1
	}
	return nil
}

func save(path string) error {
	fmt.Println("Saving:", path)

	err := os.Remove(path)
	if err != nil {
		fmt.Println("Error removing file", err)
	}

	file, err1 := os.Create(path)
	if err1 != nil {
		fmt.Println("Error creating file")
		return err1
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	err2 := encoder.Encode(&DATA)
	if err2 != nil {
		fmt.Println("Error encoding file")
		return err2
	}
	return nil
}

func main() {
	//Input a path to a gob file, opeation and if needed for the operation key, Name, Syrname and Id.
	//The input should be command line arguments
	clArgs := os.Args

	if len(clArgs) < 2 {
		fmt.Println("Input the name of a file!")
		os.Exit(1)
	}
	filePath := clArgs[1]
	load(filePath)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		commands := strings.Fields(input)

		switch len(commands) {
		case 1:
			commands = append(commands, "", "", "", "", "")
		case 2:
			commands = append(commands, "", "", "", "")
		case 3:
			commands = append(commands, "", "", "")
		case 4:
			commands = append(commands, "", "")
		case 5:
			commands = append(commands, "")
		default:
			fmt.Println("Five first arguments will be used.")
		}

		switch commands[0] {
		case "stop":
			err := save(filePath)
			if err != nil {
				fmt.Println(err)
			}
			os.Exit(0)
		case "print":
			print()
		case "add":
			if !add(commands[1], Element{commands[2], commands[3], commands[4]}) {
				fmt.Println("Add operation failed.")
			}
		case "lookup":
			element := lookup(commands[1])
			if element == nil {
				fmt.Println("There is no such key in the storage.")
			} else {
				fmt.Println(element)
			}
		case "change":
			if !change(commands[1], Element{commands[2], commands[3], commands[4]}) {
				fmt.Println("There is no such key in the storage.")
			}
		case "del":
			if !del(commands[1]) {
				fmt.Println("Delete operation failed.")
			}
		default:
			fmt.Println("Unknown command. try again.")
		}
	}
}
