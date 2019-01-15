package main

import (
	"fmt"
	"sort"
	"time"
)

//waits for value given, then appends value to provided slice
func waitThenAppend(val int, slice *[]int) {
	timerLen := time.Duration(val) * time.Millisecond
	time.Sleep(timerLen)
	*slice = append(*slice, val)
}

//runs waitThenAppend on each element of input array
func sleepSort(input []int, output *[]int) {
	for _, e := range input {
		go waitThenAppend(e, output)
	}
	for len(*output) < len(input) {
		time.Sleep(time.Nanosecond)
	}
}

//times sleepsort
func testSleepSort() {
	before := make([]int, 999)
	after := make([]int, 0)

	//full before slice with numbers
	for i := range before {
		before[i] = 999 - i
	}

	//record start time
	start := time.Now()
	//run sleep sort
	sleepSort(before, &after)
	//record end time
	end := time.Now()
	//calculate elapsed time
	elapsed := end.Sub(start)
	fmt.Println("Sleep sort: ", elapsed)
}

//times regular sort
func testRegSort() {
	before := make([]int, 999)

	//full before slice with numbers
	for i := range before {
		before[i] = 999 - i
	}

	//record start time
	start := time.Now()
	//run sort
	sort.Ints(before)
	//record end time
	end := time.Now()
	//calculate elapsed time
	elapsed := end.Sub(start)
	fmt.Println("Regular sort: ", elapsed)
}

func main() {
	testSleepSort()
	testRegSort()
}
