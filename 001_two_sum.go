package main

func twoSum(nums []int, target int) []int {
	var (
		low, high int
		lenght    = len(nums)
	)
	if lenght <= 0 {
		return nil
	}

	for low = 0; low < lenght-1; low++ {
		for high = low + 1; high < lenght; high++ {
			if nums[low]+nums[high] == target {
				return []int{low, high}
			}
		}
	}

	return nil
}
