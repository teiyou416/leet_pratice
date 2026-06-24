func minimumRefill(plants []int, capacityA int, capacityB int) int {
  times := 0    
  left , right := 0 ,len(plants) -1
  currA,currB := capacityA,capacityB
  for left < right{
	  
	  if currA < plants[left]{
             currA = capacityA
	     times++ 
	  }
	  currA -= plants[left]
	  left++
	  if currB < plants[right]{
		  currB = capacityB
             times++
	  }
	  currB -= plants[right]
	  right--
	  if left == right{
		  if currA >= currB{
			  if currA < plants[left]{
				  times++
			  }
		  }else{
			  if currB < plants[left]{
				  times++
			  }
		  }
	  }
  }
  return times
}
