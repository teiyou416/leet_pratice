func numSubarrayProductLessThanK(nums []int, k int) int {
    ans := 0
    sum := 1
    left := 0
    for right := 0; right < len(nums);right++{
         sum *= nums[right]
	 for left <= right && sum >= k{
		 sum /= nums[left]
		 left++
	 } 
	 ans += right - left + 1
    }
   return ans	
}
