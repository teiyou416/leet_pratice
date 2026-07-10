import "sort"
func solution(reservation [][]int) int{
    n := len(reservation)
    if n == 0{
	    return 0
    }
    start := make([]int,n)
    end := make([]int,n)

    for i:=0;i<n;i++{
         start = reservation[i][0]
         end = reservation[i][1]
    }

    sort.Ints(start)
    sort.Ints(end)

    room := 0
    endIdx :=0
    for i:=0;i<n;i++{
	    if start[i] > end[endIdx]{
		    room++
	    }else{
		    endIdx++
	    }
    }
    return room
}
