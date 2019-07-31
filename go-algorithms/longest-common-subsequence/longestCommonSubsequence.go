package main

import (
	"fmt"
)

func lcsLength(x string, y string) int {
	if len(x) == 0 || len(y) == 0 {
		return 0
	}
	if x[len(x)-1:] == y[len(y)-1:] {
		return 1 + lcsLength(x[0:len(x)-1], y[0:len(y)-1])
	} else {
		return max(lcsLength(x[0:len(x)-1], y[:]), lcsLength(x[:], y[0:len(y)-1]))
	}
}

func memoizedLCSLength(x string, y string, memory map[string]int) int {
	if len(x) == 0 || len(y) == 0 {
		return 0
	}

	if _, ok := memory[fmt.Sprintf("%v:%v", x, y)]; !ok {
		if x[len(x)-1:] == y[len(y)-1:] {
			memory[fmt.Sprintf("%v:%v", x, y)] = memoizedLCSLength(x[0:len(x)-1], y[0:len(y)-1], memory) + 1
		} else {
			memory[fmt.Sprintf("%v:%v", x, y)] = max(
				memoizedLCSLength(x[0:len(x)-1], y[:], memory),
				memoizedLCSLength(x[:], y[0:len(y)-1], memory))
		}
	}

	return memory[fmt.Sprintf("%v:%v", x, y)]
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
