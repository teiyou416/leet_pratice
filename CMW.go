func maxArea(height []int) int {
	MaxSum := 0
	l,r := 0, len(height)-1
	for l < r{
		currSum := min(height[l],height[r]) * (r-l)
		if currSum > MaxSum{
			MaxSum = currSum
		}
		if height[l] > height[r]{
			r--
		}else{
			l++
		}
		
	}	
	return MaxSum
}
func min(x,y int) int{
	if x > y {
		return y
	}
	return x
}
