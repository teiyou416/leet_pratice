func Solution(N int) string {
	res := make([]byte, N)

	for i := 0; i < N; i++ {
		res[i] = 'a'
	}

	if N%2 == 0 {
		res[N-1] = 'b'
	}

	return string(res)
}
