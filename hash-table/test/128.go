func longestConsecutive(nums []int) int {
	ans := 0
	hash := make(map[int]bool)
	for _, num := range nums {
		hash[num] = true
	}
	for num := range hash {
		if !hash[num-1] {
			currN := num
			currS := 1
			for hash[currN+1] {
				currN++
				currS++
			}
			if currS > ans {
				ans = currS
			}
		}
	}
	return ans
}
