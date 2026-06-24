func trap(height []int) int {
	n := len(height)	    
	premax := make([]int,n)
	sufmax := make([]int,n)
        premax[0] = height[0]	
	for i:=1;i<n;i++{
	     premax[i] = max(premax[i-1],height[i])
	}	
	sufmax[n-1] = height[n-1]
	for i:=n-2;i>=0;i--{
           sufmax[i] = max(sufmax[i+1],height[i])
	}
        water := 0	
	for i:=0;i<n;i++{
		water += min(sufmax[i],premax[i]) - height[i]
	}
	return water
}

func max(a,b int) int{
	if a > b{
		return a 
	}
	return b
}

func min(a,b int) int{
	if a > b{
		return b
	}
	return a
}
