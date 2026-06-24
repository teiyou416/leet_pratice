func trap(height []int) int{ 
    n := len(height)
    left := 0
    right := n - 1
    leftMax,rightMax := 0,0
    water := 0
    for left <= right{
	    if height[left] > leftMax{
		    leftMax = height[left]
	    } 
	    if height[right] > rightMax{
		    rightMax = height[right]
	    }
	    if leftMax < rightMax{
		    water += leftMax - height[left]
		    left++
	    }else{
		    water += rightMax - height[right]
		    right--
	    }
    }
    return water
}

