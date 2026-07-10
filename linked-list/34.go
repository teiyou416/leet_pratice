func searchRange(nums []int, target int) []int {
	left := searchBound(nums, target, true)
	right := searchBound(nums, target, false)
	return []int{left, right}

}
func searchBound(nums []int, target int, find bool) int {
	left := 0
	right := len(nums) - 1
	res := -1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			res = mid
			if find {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}

	}
	return res
}
