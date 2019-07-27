package main

import (
	"fmt"
	"os"
)

func main() {
	// an ordered array
	array := []int{4, 5, 6, 7, 10, 22, 34, 100, 107}

	// print array
	fmt.Printf("Array is: %v\n", array)

	// get num to search for
	fmt.Print("Enter num to search for: ")
	var num int
	_, err := fmt.Scanf("%d", &num)
	if err != nil {
		fmt.Printf("Input is not a valid integer\n")
		os.Exit(-1)
	}

	// search and print result
	result := binarySearch(array, num)
	fmt.Printf("Index: %d \n", result)
}
