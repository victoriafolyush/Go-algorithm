package main

import "fmt"

func getPivot(array []int) int {
	res := array[0]
	return res
}

func swapPivot(array []int, start int, end int) int {
	pivot := getPivot(array)
	pivot = array[start]
	swapIdx := start
	for i := start + 1; i <= end; i++ {
		if pivot >= array[i] {
			swapIdx++
			temp := array[i]
			array[i] = array[swapIdx]
			array[swapIdx] = temp
		}
	}
	temp := array[start]
	array[start] = array[swapIdx]
	array[swapIdx] = temp

	return swapIdx
}

func sorting(array []int, left int, right int) []int {
	if left < right {
		pivotIndex := swapPivot(array, left, right)
		sorting(array, left, pivotIndex-1)
		sorting(array, pivotIndex+1, right)
	}
	return array
}

func quickSort(array []int, getPivot func([]int) int) []int {
	sorting(array, 0, len(array)-1)
	if len(array) > 0 {
		getPivot(array)
	}
	return array
}

func main() {
	nums := []int{82, 252, 101, 83}
	fmt.Println(nums)
	fmt.Println(quickSort(nums, getPivot))
}
