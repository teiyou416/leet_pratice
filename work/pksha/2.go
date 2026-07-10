func Solution(S string) int {

	firstOneIdx := -1
	for i := 0; i < len(S); i++ {
		if S[i] == '1' {
			firstOneIdx = i
			break
		}
	}

	if firstOneIdx == -1 {
		return 0
	}

	operations := 0

	for i := firstOneIdx + 1; i < len(S); i++ {
		if S[i] == '0' {
			operations += 1
		} else {
			operations += 2
		}
	}

	operations += 1

	return operations
}
