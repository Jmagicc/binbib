package main

import "fmt"

func main() {
	fmt.Println(findLengthOfLCIS([]int{1, 3, 5, 4, 7})) // 3
	fmt.Println(findLengthOfLCIS([]int{2, 2, 2, 2, 2})) // 1

}
func findLengthOfLCIS(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return len(nums)
	}
	maxcount := 0
	slow := 0
	for slow < n-1 {
		fast := slow
		count := 1
		for fast < n-1 && nums[fast] < nums[fast+1] {
			fast++
			count++
		}
		slow++
		if count > maxcount {
			maxcount = count
		}
	}

	return maxcount
}
