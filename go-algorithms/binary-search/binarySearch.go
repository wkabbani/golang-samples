package main

func binarySearch(array []int, num int) (index int) {
	index = -1
	start := 0
	end := len(array) - 1
	for start < end {
		middle := (start + end) / 2
		switch {
		case array[middle] == num:
			index = middle
			return
		case num > array[middle]:
			start = middle + 1
		case num < array[middle]:
			end = middle - 1
		}
	}
	if start == end && array[start] == num {
		index = start
	}
	return
}
