package main

import "fmt"

func main() {
	fmt.Println(rotate([]int{55, 66, 77, 88, 99}, 1))
}

func rotate(nums []int, n int) []int {
	r := nums[n:]
	for i := 0; i < len(nums[:n]); i++ {
		r = append(r, nums[i])
	}

	return r
}
