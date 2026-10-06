package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	result := processConcurrently(x)
	fmt.Println(result)
}

const numOfGoroutines = 5
func process(val int) int{
	return val * 2
}
func processConcurrently(inVals []int) []int {
	in := make(chan int, numOfGoroutines)
	out := make(chan int, numOfGoroutines)

	for i := 0; i < numOfGoroutines; i++ {
		go func() {
			for val := range in {
				result := process(val)
				out <- result
			}
		}()
	}

	go func() {
		for _, val := range inVals {
			in <- val
		}
	}()

	outVals := make([]int, 0, len(inVals))

	for i := 0; i < len(inVals); i++ {
		outVals = append(outVals, <-out)
	}

	return outVals

}
