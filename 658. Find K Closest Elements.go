func findClosestElements(arr []int, k int, x int) []int {
    left, right := 0, len(arr)-k-1
    for left <= right {
        mid := left + (right - left)/2
        if x - arr[mid] <= arr[mid+k] - x {
            right = mid - 1
        } else {
            left = mid + 1
        }
    }
    return arr[left:left+k]
}