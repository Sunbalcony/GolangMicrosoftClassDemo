package main

import "fmt"

func aaa() [][]int {
	//	var nums = []int{0, 1, 0, 3, 13}
	//	left := 0
	//	right := 0
	//	n := len(nums)
	//	for right <  n {
	//		if nums[right] != 0 {
	//			nums[left], nums[right] = nums[right], nums[left]
	//			left++
	//		}
	//		right++
	//
	//	}
	//	fmt.Println(nums)
	//}
	var nums = []int{-1, 0, 1, 2, -1, -4}
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := i + 2; k < n; k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					fmt.Println(nums[i], nums[j], nums[k])

				}

			}
		}
	}
}
func main() {
	aaa()
}
