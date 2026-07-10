func Solution(blocks []int) int {
	n := len(blocks)
	if n <= 1 {
		return n
	}

	left := make([]int, n)
	left[0] = 0
	for i := 1; i < n; i++ {
		if blocks[i-1] >= blocks[i] {
			left[i] = left[i-1]
		} else {
			left[i] = i
		}
	}

	right := make([]int, n)
	right[n-1] = n - 1
	for i := n - 2; i >= 0; i-- {
		if blocks[i+1] >= blocks[i] {
			right[i] = right[i+1]
		} else {
			right[i] = i
		}
	}

	maxDistance := 0
	for i := 0; i < n; i++ {
		distance := right[i] - left[i] + 1
		if distance > maxDistance {
			maxDistance = distance
		}
	}

	return maxDistance
}
