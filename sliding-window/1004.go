func longestOnes(nums []int, k int) int {
	window := make(map[int]int)
	ans := 0
	left := 0
	for right:=0; right < len(nums);right++{
		inchar := nums[right]
		window[inchar]++
		for window[0] > k {
                    outchar := nums[left]
		    window[outchar]--
		    left++
		}
          currAns :=  right - left + 1
	  if currAns > ans{
		  ans = currAns
	  }
	}
	return ans
}
