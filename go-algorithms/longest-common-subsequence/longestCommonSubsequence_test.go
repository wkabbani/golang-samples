package main

import "testing"

func TestLCSLength(t *testing.T) {

	tables := []struct {
		first  string
		second string
		length int
	}{
		{"abc", "abc", 3},
		{"abc", "xyz", 0},
		{"ABCBDAB", "BDCABA", 4},
	}

	for _, table := range tables {
		result := lcsLength(table.first, table.second)
		if result != table.length {
			t.Errorf("Incorrect LCS of %v and %v, got: %d, want: %d.", table.first, table.second, result, table.length)
		}
	}
}

func TestMemoizedLCSLength(t *testing.T) {

	tables := []struct {
		first  string
		second string
		length int
	}{
		{"abc", "abc", 3},
		{"abc", "xyz", 0},
		{"ABCBDAB", "BDCABA", 4},
	}

	for _, table := range tables {
		result := memoizedLCSLength(table.first, table.second, make(map[string]int))
		if result != table.length {
			t.Errorf("Incorrect LCS of %v and %v, got: %d, want: %d.", table.first, table.second, result, table.length)
		}
	}
}
