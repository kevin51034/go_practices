package main

import "fmt"

func main() {
	var n int
	fmt.Println("Enter the number of the input:")
	fmt.Scanf("%d", &n)
	fmt.Println("Enter the inputs")
	array := make([]int, n)
	for i:=0 ; i<n; i++ {
		fmt.Scanf("%d", &array[i])
	}
	fmt.Println("The original input array is:")
	fmt.Println(array)

	sortArray := quickSort(array, 0, len(array)-1)
	fmt.Println(sortArray)
}

func quickSort(array []int,start int, end int) []int {
	if start < end {
		pivot := partition(array, start, end)
		quickSort(array, start, pivot-1)
		quickSort(array, pivot+1, end)
	}
	return array
}

func partition(array []int, start, end int) int {
	pivot := array[end]
	i := start
	for j:=start; j<end;j++ {
		if array[j] < pivot {
			array[i], array[j] = array[j], array[i]
			i++
		}
	}
	array[i], array[end] = array[end], array[i]
	return i
}