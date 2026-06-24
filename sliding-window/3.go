func lengthOfLongestSubstring(s string) int {
	window := make(map[byte]int) 
	left := 0
	ans := 0
	for right:=0;right<len(s);right++{
		inchar := s[right]
		window[inchar]++
		for window[inchar] > 1{
			outchar := s[left]
			window[outchar]--
			left++
		}  
		currAns := right - left + 1
		if currAns > ans{
			ans = currAns
		}
		
	}
	return ans
}
