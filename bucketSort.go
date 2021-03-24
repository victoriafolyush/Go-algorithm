package main

import "fmt"

func isEqual(a []int) bool {
	for i := 1; i < len(a); i++ {
		if a[i] != a[0] {
			return false
		}
	}
	return true
}

func msBits(array []int, element int) int {
	max, min := array[0], array[0]
	for i := 1; i <= len(array)-1; i++ {
		if array[i] > max {
			max = array[i]
		} else if array[i] < min {
			min = array[i]
		}
	}
	res := int((float64(element-min) / float64(max-min+1)) * float64(len(array)))
	return res
}

func bucketSort(array []int, msBits func([]int, int) int) []int {
	buckets := make([][]int, 0)
	buckets = make([][]int, len(array))

	for i := 0; i < len(array); i++ {
		bucketNum := msBits(array, array[i])
		buckets[bucketNum] = append(buckets[bucketNum], array[i])
	}
	for i := 0; i < len(buckets); i++ {

		if len(buckets[i]) > 1 && isEqual(buckets[i]) == false {
			buckets[i] = bucketSort(buckets[i], msBits)
		}
	}

	k := 0
	for i := 0; i < len(buckets); i++ {
		for j := 0; j < len(buckets[i]); j++ {
			array[k] = buckets[i][j]
			k += 1
		}
	}

	return array
}

func main() {
	nums := []int{9, 9,  0 ,252, -7, 252, 9}
	fmt.Println(bucketSort(nums, msBits))

}
