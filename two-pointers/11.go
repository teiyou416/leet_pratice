func maxArea(height []int) int {
        maxAns := 0	   
        left := 0	
	right := len(height) - 1
	for left < right{
           currAns := min(height[left],height[right]) * (right-left) 
	   if currAns > maxAns{
		   maxAns = currAns
	   }
	   if height[left] < height[right]{
		   left++
	   }else{
		   right--
	   }
	}
	return maxAns
}
func min(a,b int) int{
	if a > b {
		return b
	}
	return a
}
