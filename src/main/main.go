package main

import (
	"fmt"
	"sortMethods"
	"runtime"
	"reflect"
	"time"
)

func main() {
	listLength := *readListLength()

	order := readOrderMethod();

	list, err := generateList(listLength, order)
	if err != nil {
		fmt.Printf("Error generating list (%v)", err)
		return
	}

	println("List to order: ", list)

	sortFunctions := []sortMethods.SortFunc{
		sortMethods.SelectSort,
		sortMethods.InsertSort,
		sortMethods.QuickSort,
	}

	resultChan := make(chan string)

	for _, sortFunc := range sortFunctions {
		go benchmark(sortFunc, list, resultChan)
	}

	for i := 0; i < len(sortFunctions); i++ {
		println()
		println(<-resultChan)
	}
}

func readListLength() *uint32 {
	var numberOfElements uint32
	for {
		fmt.Println("Size of list elements:")
		_, err := fmt.Scan(&numberOfElements)
		if err == nil {
			return &numberOfElements
		} else {
			fmt.Printf("Error reading input (%v), please try again\r\n", err)
		}
	}
}

func readOrderMethod() ListOrder {
	var orderMethodSelection uint
	for {
		println("Choose the order of the elements in the list:")
		println("1) Ordered (1,2,3...)")
		println("2) Inverse (3,2,1...)")
		println("3) Random (2,3,1...)")
		_, err := fmt.Scan(&orderMethodSelection)
		if err == nil {
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

func benchmark(sortFunc sortMethods.SortFunc, list *[]uint32, resultChan chan string) {
	listCopy := make([]uint32, len(*list))
	copy(listCopy, *list)

	startTime := time.Now()
	checkCount, swapCount := sortFunc(listCopy)
	elapsed := time.Since(startTime)

	var result string
	result += fmt.Sprintln(GetFunctionName(sortFunc))
	result += fmt.Sprintln("Duration: ", elapsed)
	result += fmt.Sprintln("Checks: ", checkCount)
	result += fmt.Sprintln("Swaps: ", swapCount)

	resultChan <- result
}

func GetFunctionName(function interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(function).Pointer()).Name()
}