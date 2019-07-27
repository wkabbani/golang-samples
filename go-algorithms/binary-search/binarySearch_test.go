package main

import "testing"

func TestBinarySearch(t *testing.T) {

	array := []int{4, 5, 6, 7, 10, 22, 34, 100, 107}

	tables := []struct {
		input []int
		num   int
		index int
	}{
		{array, 107, 8},
		{array, 4, 0},
		{array, 34, 6},
		{array, 10, 4},
		{array, 200, -1},
		{array, 1, -1},
		{array, 0, -1},
		{array, -50, -1},
	}

	for _, table := range tables {
		result := binarySearch(table.input, table.num)
		if result != table.index {
			t.Errorf("Search of %d in %v was incorrect, got: %d, want: %d.", table.num, table.input, result, table.index)
		}
	}
}
