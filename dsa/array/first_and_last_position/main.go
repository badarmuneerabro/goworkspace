package main

import "fmt"

func findFirstAndLastPosition(nums[] int, target int) []int{
	lowerBound := lowerBound(nums, target)
	
	if lowerBound == len(nums) || nums[lowerBound] != target {
		return []int{-1, -1}
	}
	
	/*if upperBound - 1 == len(nums) || nums[upperBound - 1] != target {
		upperBound = -1
	}*/
	
	return []int{lowerBound, upperBound(nums, target)  - 1}
}

func lowerBound(nums[] int, target int) int{
	start := 0
	end := len(nums) - 1
	ans := len(nums) /*Covers where all the values < target*/
	
	for(start <= end){
		mid := (start + end) / 2
		if nums[mid] >= target {
			ans = mid
			end = mid - 1
		}else{
			start = mid + 1
		}
	}
	
	return ans
}

func upperBound(nums[] int, target int) int{
	start := 0
	end := len(nums) - 1
	ans := len(nums)
	
	for (start <= end){
		mid := (start + end) / 2
		if nums[mid] > target{
			ans = mid
			end = mid - 1
		}else{
			start = mid + 1
		}
	}
	
	return ans;
}

func main(){
	//a := []int{5,7,7,8,8,10}
	a := []int{1}
	result := findFirstAndLastPosition(a, 1)
	fmt.Println(result)
}