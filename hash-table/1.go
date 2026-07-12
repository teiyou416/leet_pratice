func twoSum(nums []int, target int) []int {
	hash := make(map[int]int)
	for i, num := range nums {
		diff := target - num
		if val, ok := hash[diff]; ok {
			return []int{val, i}
		}
		hash[num] = i
	}
	return nil
}
