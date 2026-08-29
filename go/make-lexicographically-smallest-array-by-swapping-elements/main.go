package main

import (
	"fmt"
	"slices"
)

func lexicographicallySmallestArray(nums []int, limit int) []int {
	gotIt := false
	for !gotIt {
		for i := 0; i < len(nums)-1; i++ {
			min := slices.Min(nums[i : ])
			minIndex := slices.Index(nums[i:], min) + i
			if minIndex != i && minIndex-i <= limit && nums[i] - min < limit {
				nums[i], nums[minIndex] = nums[minIndex], nums[i]
			}
	}
	return nums
}

func main() {
	nums := []int{1, 7, 6, 18, 2, 1}
	limit := 3
	result := lexicographicallySmallestArray(nums, limit)
	fmt.Println(result)
}
