package main

import (
	"fmt"
	"sort"
	"time"
)

func waitThenAppend(val int, slice *[]int) {
	timerLen := time.Duration(val) * time.Millisecond
	time.Sleep(timerLen)
	*slice = append(*slice, val)
}

func sleepSort(input []int, output *[]int) {
	for _, e := range input {
		go waitThenAppend(e, output)
	}
	for len(*output) < len(input) {
		time.Sleep(time.Nanosecond)
	}
}

func testSleepSort() {
	before := make([]int, 999)
	after := make([]int, 0)

	for i := range before {
		before[i] = 999 - i
	}

	start := time.Now()
	sleepSort(before, &after)
	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Println("Sleep sort: ", elapsed)
}

func testRegSort() {
	before := make([]int, 999)

	for i := range before {
		before[i] = 999 - i
	}

	start := time.Now()
	sort.Ints(before)
	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Println("Regular sort: ", elapsed)
}

func main() {
	testSleepSort()
	testRegSort()
}
