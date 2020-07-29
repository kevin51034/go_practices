package main

import "fmt"

func main() {
	var n int
	fmt.Println("Enter the number of the inputs")
	fmt.Scanf("%d", &n)
	fmt.Println("Enter the inputs")
	array := make([]int, n)
	for i:=0; i<n; i++ {
		fmt.Scanf("%d", &array[i])
	}
	fmt.Println("The original array is:")
	fmt.Println(array)

	sortArray := heapSort(array)
	fmt.Println(sortArray)
}
