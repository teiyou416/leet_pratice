func ans(nums []int) []int{
	n := len(nums)
	ans := make([]int,n)
	
	left := make([]int,n)
	right := make([]int,n)

	stack := make([]int,0,n)

	for i,num := range nums{
		for len(stack)>0 && num >= nums[stack[len(stack)-1]]{
		        top := len(stack) - 1
			idx := stack[top]
                        right[idx] = i
			stack = stack[:top]
		}
		stack = append(stack,i)
	}

	stack = stack[:0]

	for i:=n-1;i>=0;i--{
		for len(stack)>0 && nums[i] >= nums[stack[len(stack)-1]]{
			top := len(stack) -1
			idx := stack[top]
                        left[idx] = i
			stack = stack[:top]
		}
		stack = append(stack,i)
	}
	for i:=0;i< n;i++{
		ans[i] = (i-left[i]) + (right[i] - i) -1
	}
	return ans
}
