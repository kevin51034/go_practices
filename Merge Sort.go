package main

import "fmt"

func main() {
	var n int
	//var array []int
	fmt.Println("Enter the number of the input:")
	fmt.Scanf("%d", &n)
	fmt.Println("Enter the inputs")
	array := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &array[i])
	}
	fmt.Println("The original input array is:")
	fmt.Println(array)

	fmt.Println("The sorted array is:")
	sortArray := mergeSort(array)
	fmt.Println(sortArray)
}

func mergeSort(array []int) []int {
	if len(array) <= 1 {
		return array
	}
	midPoint := len(array) / 2

	left := mergeSort(array[:midPoint])
	right := mergeSort(array[midPoint:])

	return mergeTwoArray(left, right)
}

func mergeTwoArray(left, right []int) []int {

	i := 0
	j := 0
	result := make([]int, 0)
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// use three dot notation to concatenate two array
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}
