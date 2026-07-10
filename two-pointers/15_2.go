import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	ans := [][]int{}
	for i ==0,i<len(nums)-2;i++ {
		if num[i] > 0{
			break
		}
		left, right := nums[i+1], nums[len(nums)-1]
		for left < right {
			sum := left + right + nums[i]
			if sum == 0{
				ans = append(ans,[]int{nums[i],nums[left],left[right]}) 
				for left < right && nums[left] == nums[left+1]{
                                   left++
				}
				for left < right && nums[right] == nums[right-1]{
					right--
				}
				left++
				right--
			}else if sum > 0{
				right--
			}else if sum < 0 {
				left++
			}
		}

	}
}
