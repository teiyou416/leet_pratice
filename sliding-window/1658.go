import "math"
func minOperations(nums []int, x int) int {
     totalSum := 0
     for _,num := range nums{
	     totalSum += num
     }
     target := totalSum - x
     left := 0
     ans := math.MinInt32
     for right:=0;right<len(nums);right++{
	     
     }
}

