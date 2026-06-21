func longestConsecutive(nums []int) int {
          numSet := map[int]bool{}	    
	  for _, num := range nums{
		  numSet[num] = true
	  }	
	  LongestS := 0
	  for num := range numSet{
		  if !numSet[num - 1]{
                     currN := num
		     currS := 1
		  
		  for numSet[currN+1]{
			  currN++ 
			  currS++
		  }
		  if currS > LongestS{
			  LongestS = currS
		  }
	  }
  }
	  return LongestS
}
