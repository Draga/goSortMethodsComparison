package main

import (
	"fmt"
)

func main() {
	elementsNumString := readNumberOfElements()
	fmt.Print(elementsNumString)
}

func readNumberOfElements() (numberOfElements int) {
	for {
		fmt.Println("Number of list elements: ")
		_, err := fmt.Scan(&numberOfElements)
		if err == nil{
			return
		} else {
			fmt.Printf("Error reading input (%v), please try again\r\n", err)
		}
	}
}