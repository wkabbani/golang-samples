package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	fmt.Printf("Enter first string sequence:\n")
	firstSequence := readLine(reader)

	fmt.Printf("Enter second string sequence:\n")
	secondSequence := readLine(reader)

	result1 := lcsLength(firstSequence, secondSequence)

	result2 := memoizedLCSLength(firstSequence, secondSequence, make(map[string]int))

	fmt.Printf("LCS Length (Simple): %d \n", result1)
	fmt.Printf("LCS Length (Memoized): %d \n", result2)

}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}
