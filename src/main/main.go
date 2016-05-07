package main

import (
	"fmt"
	"sortMethods"
)

func main() {
	listLength := *readListLength()
	fmt.Print(listLength)

	order := readOrderMethod();
	fmt.Println(order)

	list, err := generateList(listLength, order)
	fmt.Println(list, err)
	if err != nil{
		fmt.Printf("Error generating list (%v)", err)
		return
	}

	sortedList, checkCount ,swapCount := sortMethods.SelectSort(*list)

	fmt.Println(sortedList)
	fmt.Println(checkCount)
	fmt.Println(swapCount)
}

func readListLength() *uint32 {
	var numberOfElements uint32
	for {
		fmt.Println("Size of list elements:")
		_, err := fmt.Scan(&numberOfElements)
		if err == nil{
			return &numberOfElements
		} else {
			fmt.Printf("Error reading input (%v), please try again\r\n", err)
		}
	}
}

func readOrderMethod() ListOrder {
	var orderMethodSelection uint
	for {
		fmt.Println("Choose the order of the elements in the list:")
		fmt.Println("1) Ordered (1,2,3...)")
		fmt.Println("2) Inverse (3,2,1...)")
		fmt.Println("3) Random (2,3,1...)")
		_, err := fmt.Scan(&orderMethodSelection)
		if err == nil{
			switch orderMethodSelection {
			case 1:
				return ordered
			case 2:
				return inverse
			case 3:
				return random
			default:
				fmt.Printf("%v is not a valid choice!\r\n", orderMethodSelection)
			}
		} else {
			fmt.Printf("Error reading input (%v), please try again\r\n", err)
		}
	}
}