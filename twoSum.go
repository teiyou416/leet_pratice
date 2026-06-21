func twoSum(nums []int, target int) []int {
	// hash := map[int]int{}
	hash := make(map[int]int)
	for i, num := range nums{
		if  p ,ok := hash[target-num];ok{
			return []int{p,i}
		}
		hash[num] = i 
	}  	
	return nil
}

