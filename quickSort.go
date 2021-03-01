package main

import "fmt"

func getPivot(array []int) int {
	res := array[0]
	return res
}

func quickSort(array []int, getPivot func ([] int) int) []int {
	left := 0
	right := len(array)-1

	if left < right {
		pivot := getPivot(array)
		swapIdx := left
		for i := left + 1; i <= right; i++ {
			if pivot >= array[i] {
				swapIdx++
				array[i], array[swapIdx] = array[swapIdx], array[i]
			}
		}
		array[left], array[swapIdx] = array[swapIdx], array[left]

		quickSort(array[:swapIdx], getPivot)
		quickSort(array[swapIdx + 1:], getPivot)
	}
	return array
}

func main() {
	nums := []int{ 7, 7, -7 ,252 ,7 ,-95 ,15 }
	fmt.Println(nums[: 7])
	fmt.Println(quickSort(nums, getPivot))
}
