package main

import (
	"slices"
	"testing"
)

func TestSortIntegers(t *testing.T) {
	nums := []int{1, 3, 2, 4}

	result := make([]int, len(nums))

	copy(result, nums)

	slices.Sort(nums)
	SortIntegers(result)

	if !slices.Equal(result, nums) {
		t.Errorf("Error")
	}
}

// func main() {
// 	nums := []int{1, 3, 2, 4}

// 	result := make([]int, len(nums))

// 	copy(result, nums)

// 	sort.Ints(nums)

// 	fmt.Println(nums, result)
// }
