package main

import "fmt"

func NumberOfOccurences(nums []int, target int) int {
	lowerBound := lowerBound(nums, target)
	fmt.Println(lowerBound)
	upperBound := upperBound(nums, target)
	fmt.Println(upperBound)

	return upperBound - lowerBound
}

func lowerBound(nums []int, target int) int {
	start := 0
	end := len(nums)
	ans := -1
	for start <= end {
		fmt.Println("Working...", start, end)
		mid := (start + end) / 2
		if nums[mid] >= target {
			ans = mid
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return ans
}

func upperBound(nums []int, target int) int {
	start := 0
	end := len(nums)
	ans := -1

	for start <= end {
		mid := (start + end) / 2
		if nums[mid] > target {
			ans = mid
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	
	return ans
}

func main() {
	a := []int{10, 20, 20, 20, 20, 20, 30, 40, 50}
	target := 20
	n := NumberOfOccurences(a, target)
	fmt.Printf("Number of occurrences of %d, are %d", target, n)
}
