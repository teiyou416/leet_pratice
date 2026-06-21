func findAnagrams(s string, p string) []int {
	// var window []int
	window := make(map[byte]int)
	var ans []int
	// for _,word:= range p{
	//         window[word]++
	// }
	for i:=0 ; i<len(p);i++{
		window[p[i]]++
	}
	L := 0	

	diffCount := len(window)
	for R:=0;R<len(s);R++{
		in := s[R]
		if _,ok := window[in];ok{
			window[in]--
			if window[in]==0{
				diffCount--
			}
		}
		if R-L+1 > len(p){
			out := s[L]
			if _,ok := window[out];ok{
				if window[out]==0{
					diffCount++
				}
				window[out]++
			}
			L++
		}
		if diffCount == 0{
			ans = append(ans,L)
		}

	}
	return ans
}
