// Merge Sort
func sortArray(nums []int) []int {
    if len(nums) <= 1 {
        return nums
    }
    left := sortArray(nums[:len(nums)/2])
    right := sortArray(nums[len(nums)/2:])
    return mergeTwoList(left, right)
}

func mergeTwoList(l1, l2 []int) []int {
    list := make([]int, 0)
    i, j := 0, 0
    for i<len(l1) && j<len(l2) {
        if l1[i] <= l2[j] {
            list = append(list, l1[i])
            i++
        } else {
            list = append(list, l2[j])
            j++
        }
    }
    list = append(list, l1[i:]...)
    list = append(list, l2[j:]...)
    return list
}

// Quick Sort
func sortArray(nums []int) []int {
    return quickSort(nums, 0, len(nums)-1)
}

func quickSort(nums []int, start, end int) []int {
    if start < end {
        pivot := partition(nums, start, end)
        quickSort(nums, start, pivot-1)
        quickSort(nums, pivot+1, end)
    }
    return nums
}

func partition(nums []int, start, end int) int {
    pivot := nums[end]
    j := start
    for i:=start; i<end; i++ {
        if nums[i] < pivot {
            nums[i], nums[j] = nums[j], nums[i]
            j++
        }
    }
    nums[end], nums[j] = nums[j], nums[end]
    return j
}

// Heap Sort
func sortArray(nums []int) []int {
    heapSort(nums)
    return nums
}

func heapSort(nums []int) {
    for i:=len(nums)/2-1; i>=0; i-- {
        maxHeapify(i, len(nums), nums)
    }
    for i:=len(nums)-1; i>=0; i-- {
        nums[0], nums[i] = nums[i], nums[0]
        maxHeapify(0, i, nums)
    }
}

func maxHeapify(index, end int, nums []int) {
    for {
        i := index
        l := i * 2 + 1
        r := i * 2 + 2
        if l<end && nums[i] < nums[l] {
            i = l
        }
        if r<end && nums[i] < nums[r] {
            i = r
        }
        if i==index {
            return
        }
        nums[i], nums[index] = nums[index], nums[i]
        index = i
    }
}