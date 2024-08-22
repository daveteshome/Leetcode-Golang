package main

import "fmt"

func main() {
	fmt.Println(containsDuplicate([]int{1, 2, 3, 1}))
	fmt.Println(containsDuplicate([]int{1, 2, 3, 4}))
	fmt.Println(containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}))
}

func containsDuplicate(nums []int) bool {
	freq := make(map[int]bool)

	for i := 0; i < len(nums); i++ {
		if _, ok := freq[nums[i]]; ok {
			return true
		}

		freq[nums[i]] = true
	}

	return false
}
