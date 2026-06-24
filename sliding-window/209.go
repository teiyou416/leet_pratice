import "math"
func minSubArrayLen(target int, nums []int) int {
       n := len(nums)	    
       left := 0
       minAns := math.MaxInt32
       sum := 0 
       for right:=0;right<n;right++{
             sum += nums[right]
	     for sum >= target{
		     currAns := right - left + 1
	     
	     if currAns < minAns{
                 minAns = currAns
	     }
	     sum -= nums[left]
	     left++
       }
       }
       if minAns == math.MaxInt32{
	       return 0
       }
       return minAns
       
}
