package main

func lexicographicallySmallestArray(nums []int, limit int) []int {
	for limit > 0 {
		for i := 0; i < len(nums)-1; i++ {
			if nums[i] > nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
				limit--
				if limit == 0 {
					break
				}
			}
		}
	}
	return nums
}
