package main

import "fmt"


func main() {
	var n int
	//var array []int
	fmt.Println("Enter the number of the input:")
	fmt.Scanf("%d", &n)
	fmt.Println("Enter the inputs")
	array := make([]int, n)
	for i:=0 ; i<n; i++ {
		fmt.Scanf("%d", &array[i])
	}
	fmt.Println("The original input array is:")
	fmt.Println(array)

	fmt.Println("The sorted array is:")
	sortArray := heapSort(array)
	fmt.Println(sortArray)
}

func heapSort(array []int) []int {
	// build maxHeap
	for i:=len(array)/2-1; i>=0; i-- {
		sink(array, i, len(array))
	}

	// inverse the maxHeap to sorted array
	// swap fist and last node then pretend the last node is no longer exist(cuz it's already sorted)
	for i:=len(array)-1; i>=0; i-- {
		array[i], array[0] = array[0], array[i]
		sink(array, 0, i)
	}
	return array
}

func sink(array []int, i, length int) {
	for {
		l := i * 2 + 1
		r := i * 2 + 2
		index := i
	
		if l < length && array[l] > array[index] {
			index = l
		}
		if r < length && array[r] > array[index] {
			index = r
		}
	
		if index == i {
			return
		}
	
		array[index], array[i] = array[i], array[index]
	
		i = index
	}

}