package main
import "fmt"

func main(){
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 8, 10, 11}
	result := lowerBound(a, 9)
	
	fmt.Println(result)
}

//lowerBound = smallest element that is >= target
func lowerBound(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	mid := 0
	result := len(nums)
	
	for left < right{
		mid = (left + right) / 2
		if nums[mid] >= target{
			result = mid
			right = mid - 1
		}else{
			left = mid + 1
		}
	}
	
	return result
}