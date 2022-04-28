package main

import "fmt"

func Test_004() {
	var nums1 = []int{}
	var nums2 = []int{2, 3}
	fmt.Println(findMedianSortedArrays2(nums1, nums2))
	// nums1 = []int{}
	// nums2 = []int{2, 4, 6}
	// fmt.Println(findMedianSortedArrays2(nums1, nums2))
}

func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	var length = len(nums1) + len(nums2)
	var sum int

	if len(nums1) == 0 || len(nums2) == 0 {
		if len(nums1) != 0 {
			sum = nums1[(length)/2] + nums1[(length-1)/2]
		} else {
			sum = nums2[(length)/2] + nums2[(length-1)/2]
		}
		return float64(sum) / 2
	}

	var v1, v2 int
	var idx1, idx2 int
	for i := 0; i < (length+1)/2; i++ {
		if idx1+1 > len(nums1) {
			v1 = nums2[idx2]
			idx2 = idx2 + 1
			continue
		} else if idx2+1 > len(nums2) {
			v1 = nums1[idx1]
			idx1 = idx1 + 1
			continue
		}

		if nums1[idx1] < nums2[idx2] {
			v1 = nums1[idx1]
			idx1 = idx1 + 1
		} else if nums1[idx1] > nums2[idx2] {
			v1 = nums2[idx2]
			idx2 = idx2 + 1
		} else if nums1[idx1] == nums2[idx2] {
			v1 = nums1[idx1]
			idx1 = idx1 + 1
		}
	}

	idx1, idx2 = 0, 0
	for i := 0; i < (length+2)/2; i++ {
		if idx1+1 > len(nums1) {
			v2 = nums2[idx2]
			idx2 = idx2 + 1
			continue
		} else if idx2+1 > len(nums2) {
			v2 = nums1[idx1]
			idx1 = idx1 + 1
			continue
		}
		if nums1[idx1] < nums2[idx2] {
			v2 = nums1[idx1]
			idx1 = idx1 + 1
		} else if nums1[idx1] > nums2[idx2] {
			v2 = nums2[idx2]
			idx2 = idx2 + 1
		} else if nums1[idx1] == nums2[idx2] {
			v2 = nums1[idx1]
			idx1 = idx1 + 1
		}
	}
	return float64(v1+v2) / 2
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var (
		nums       = make([]int, 0)
		idx1, idx2 int
	)

	if len(nums1) == 0 {
		nums = nums2
	} else if len(nums2) == 0 {
		nums = nums1
	} else {
		for {
			if nums1[idx1] > nums2[idx2] {
				nums = append(nums, nums2[idx2])
				idx2 = idx2 + 1
			} else if nums1[idx1] < nums2[idx2] {
				nums = append(nums, nums1[idx1])
				idx1 = idx1 + 1
			} else {
				nums = append(nums, nums2[idx2])
				nums = append(nums, nums1[idx1])
				idx1 = idx1 + 1
				idx2 = idx2 + 1
			}
			if idx1 >= len(nums1) {
				nums = append(nums, nums2[idx2:]...)
				break
			}
			if idx2 >= len(nums2) {
				nums = append(nums, nums1[idx1:]...)
				break
			}
		}
	}

	i, j := 0, len(nums)-1
	for {
		if j > i {
			i++
			j--
		} else if j == i {
			return float64(nums[j])
		} else if j < i {
			return float64((nums[i] + nums[j])) / 2
		}
	}
}
