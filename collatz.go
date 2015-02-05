package main

import (
	"fmt"
	"time"
)

func collatz(n int) int {
	if n%2 == 0 {
		return n / 2
	}
	return n*3 + 1
}

func collatzSeq(start int) []int {
	n := start
	ls := []int{}
	for n != 1 {
		n = collatz(n)
		ls = append(ls, n)
		time.Sleep(1e7)
	}
	return append(ls, 1)

}

func main() {
	for i := 1; i < 100; i++ {
		fmt.Println(collatzSeq(i))
	}
}
