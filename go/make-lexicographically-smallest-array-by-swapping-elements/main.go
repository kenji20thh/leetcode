package main

import (
	"fmt"
	"slices"
)

func lexicographicallySmallestArray(nums []int, limit int) []int {
	gotIt := false
	for !gotIt {
		gotIt = true
		for i := 0; i < len(nums)-1; i++ {
			min := slices.Min(nums[i:])
			minIndex := slices.Index(nums[i:], min) + i
			if minIndex != i && Abs(nums[i], min) <= limit {
				nums[i], nums[minIndex] = nums[minIndex], nums[i]
				gotIt = false
			}
			fmt.Println(nums)
		}
	}
	return nums
}

func Abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func main() {
	nums := []int{1, 7, 6, 18, 1, 2}
	limit := 3
	result := lexicographicallySmallestArray(nums, limit)
	fmt.Println(result) // Output: [1, 3, 2, 4, 5]
}
