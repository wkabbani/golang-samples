package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	fmt.Println("Enter N:")
	boardSize := readLine(reader)

	n, err := strconv.Atoi(boardSize)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("N is %d\n", n)

	board := makeBoard(n)

	nQueen(board, 0)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}
	return strings.TrimRight(string(str), "\r\n")
}
