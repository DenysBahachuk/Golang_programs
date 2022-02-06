package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	// Enter the IPv4 adress to check as a command line argument
	input := os.Args

	if len(input) != 2 {
		fmt.Println("Enter one IPv4 adress to check!")
		return
	}

	checkIP := input[1]
	dots := strings.Count(checkIP, ".")

	if dots != 3 {
		fmt.Println("The IP is incorrect! " +
			"There must be three dots between four fields with numbers in the IPv4 address!")
		return
	}

	partsIP := strings.Split(checkIP, ".")

	for i, v := range partsIP {
		partIP, err := strconv.Atoi(v)
		if err != nil {
			fmt.Printf("The %d part %s in the IP is incorrect. "+
				"All the parts must be numbers.\n%v", i+1, v, err)
			return
		}
		if partIP < 0 || partIP > 255 {
			fmt.Printf("The %d part %s in the IP is incorrect. "+
				"The IP parts must be in range between 0 and 255.\n", i+1, v)
			return
		}
	}
	fmt.Printf("IP address %s is correct!", checkIP)
}
