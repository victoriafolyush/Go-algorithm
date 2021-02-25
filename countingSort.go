package main

import (
	"fmt"
)

func countingSort (array []int) []int {
	if len(array) == 0 {
		return array
	}
	var maxVal, minVal = array[0], array[0]
	for i := 0; i < len(array); i++ {
		if maxVal < array[i] {
			maxVal = array[i]
		}
		if minVal > array[i] {
			minVal = array[i]
		}
	}
	counts := make([]int, maxVal-minVal+1)
	for i:= 0; i < len(array); i++{
		counts[array[i] - minVal]++
	}
	k := 0
	for i:= 0; i < len(counts); i++ {
		currentVal := i + minVal
		for j := 0; j < counts[i]; j++{
			array[k] = currentVal
			k++
		}
	}
	return array
}

func main() {
	var nums = []int {-2,3,4,-2,4,1,7,7}

	fmt.Println(nums)
	fmt.Println(countingSort(nums))
}
