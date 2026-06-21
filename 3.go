func lengthOfLongestSubstring(s string) int {
       // window := make([3]int)
       window := make(map[byte]int)
       // var window [3]int
       left := 0
       ans := 0
	
       for right := 0; right < len(s);right++{
            in := s[right]
	    window[in]++
	    for window[in]>1{
		    out := s[left]
		    window[out]--
		    left++
	    }
	    if right - left +1 > ans{
		    ans = right - left + 1
	    }
       }
       return ans
}
