// Easy
// https://leetcode.com/problems/search-insert-position/

package main

import "fmt"

func searchInsert(nums []int, target int) int {
	low := 0
	high := len(nums) - 1
	var mid int

	for low <= high {
		mid = (low + high) / 2
		midVal := nums[mid]

		if midVal == target {
			return mid
		} else if midVal > target {
			high = mid - 1
		} else if midVal < target {
			low = mid + 1
		}
	}

	for {
		if high < 0 {
			return 0
		} else if nums[high] > target {
			return high
		} else {
			return high + 1
		}
	}
}

func main() {
	nums := []int{1, 3}

	fmt.Println(searchInsert(nums, 3))
}

