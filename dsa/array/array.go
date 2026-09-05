package main
import "fmt"
import "slices"
import "math"
func main(){
	a := []int{1,8,6,2,5,4,8,3,7}
	//a := []int{1, 1}
	n := maxArea(a)
	fmt.Println(n)
}

func removeElement(nums []int, val int) int{
	if len(nums) == 0 {
		return 0;
	}
	
	index := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[index] = nums[i]
			index++
		}
	}
	
	return index
}

func removeDuplicatesFromSortedArray(nums []int)int{
	if len(nums) == 0 {
		return 0
	}
	
	index := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[index] {
			index++
			nums[index] = nums[i]
		}
	}
	
	return index + 1
}

func removeDuplicates(nums []int)int{
	var index int = 1
	for i := 1; i < len(nums); i++ {
		if nums[i - 1] != nums[i] {
			nums[index] = nums[i]
			index++
		}
	}
	
	return index
}


func threeSum(nums []int) [][]int {
    slices.Sort(nums)
    result := [][]int{}

    for i := 0; i < len(nums); i++ {
        left := i + 1;
        right := len(nums) - 1
		
		if i > 0 && nums[i] == nums[i - 1]{
			continue
		}

        for left < right {
            sum := nums[i] + nums[left] + nums[right]
            if sum == 0 {
                result = append(result, []int{nums[i], nums[left], nums[right]})
				
				for left < right && nums[left] == nums[left + 1]{
					left++
				}
				for left < right && nums[right] == nums[right - 1] {
					right--
				}
            }else if sum < 0{
                left++
            }else{
                right--
            }
        }
    }

    return result
}

func maxArea(height []int) int {
    left := 0
	right := len(height) - 1
	max := 0
	
	for left < right {
		h := int(math.Min(float64(height[left]), float64(height[right])))
		width := right - left;
		
		area := width * h
		
		if area > max {
			max = area
		}
		
		if height[left] < height[right]{
			left++
		}else{
			right--
		}
	}
	
	return max
}