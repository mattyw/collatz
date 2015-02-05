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

func collatzSeqOneness(oneness map[int]bool, n int) bool {
	_, ok := oneness[n]
	return ok
}

func main() {
	oneness := map[int]bool{}
	for i := 1; i < 100; i++ {
		fmt.Printf("trying %v\n", i)
		if collatzSeqOneness(oneness, i) {
			fmt.Printf("skipping %v, we know it has oneness\n", i)
			continue
		}
		ls := collatzSeq(i)
		for _, i := range ls {
			oneness[i] = true // For every item in the seq we ended up at 1, therefore we know all the items in this seq have oneness.
		}
	}
	for k, _ := range oneness {
		fmt.Printf("%v, ", k)
	}
	fmt.Println("")
}
