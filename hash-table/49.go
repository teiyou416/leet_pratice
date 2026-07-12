func groupAnagrams(strs []string) [][]string {
	hash := make(map[[26]int]string)
	for _, str := range strs {
		cnt := [26]int{}
		for _, b := range str {
			cnt[b-'a']++
		}
		hash[cnt] = append(hash[cnt], str)
	}
	ans := make([][]string, 0, len(hash))
	for _, v := range hash {
		ans = append(ans, v)
	}
	return ans
}
