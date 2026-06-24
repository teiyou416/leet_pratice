func processStr(s string) string {
	var result []rure	    
	for _,ch := range s {
		switch ch{
		case '*':
                      result = result[:len(result)-1]
		case '#':
		     copypart = make([]rure,len(result))
		     copy(copypart,result)
		     result = append(result,copypart...)
		case '%':
			for i,j := 0,len(result)-1;i<j;i,j = i+1,j-1{
				result[i],result[j] = result[j],result[i]
			} 
		default:
			result = append(result,ch)
		}
	}
	return string(result)
	
}
