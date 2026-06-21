func subarraySum(nums []int, k int) int {
     presumMap := make(map[int]int)	    
     presumMap[0] = 1 	

     count := 0
     currentSum := 0

     for i:=0;i<len(nums);i++{
         currentSum += nums[i]	     
	 target := currentSum - k
	 if times,exist := presumMap[target];exist{
		 count += times
	 }
	 presumMap[currentSum]++
     }
     return count
}
