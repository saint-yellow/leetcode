// LeetCode 主站 Problem Nr. 88: 合并两个有序数组
package main

import (
	"fmt"
	"sort"
)

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3

	fmt.Println("before mergence:")
	fmt.Println("nums1 = ", nums1)
	fmt.Println("nums2 = ", nums2)

	merge(nums1, m, nums2, n)

	fmt.Println("after mergence:")
	fmt.Println("nums1 = ", nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	method2(nums1, m, nums2, n)
}

func method1(nums1 []int, m int, nums2 []int, n int) {
	copy(nums1[m:], nums2)
	sort.Ints(nums1)
}

func method2(nums1 []int, m int, nums2 []int, n int) {
	nums3 := make([]int, m+n)
	i, j, k := 0, 0, 0

	for i < m && j < n {
		if nums1[i] < nums2[j] {
			nums3[k] = nums1[i]
			i += 1
			k += 1
		} else if nums1[i] > nums2[j] {
			nums3[k] = nums2[j]
			j += 1
			k += 1
		} else {
			nums3[k] = nums1[i]
			i += 1
			k += 1

			nums3[k] = nums2[j]
			j += 1
			k += 1
		}
	}

	if i < m {
		copy(nums3[k:], nums1[i:])
	}

	if j < n {
		copy(nums3[k:], nums2[j:])
	}

	copy(nums1, nums3)
}
