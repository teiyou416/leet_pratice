moveZeroes(nums []int)  {
       zeroCount := 0
       for i,num := range nums{
	       if num != 0{
                    nums[i],nums[i-zeroCount] = nums[i-zeroCount], nums[i]
	       }
	       zeroCount++
       } 
	
}
