func isPerfectSquare(num int) bool {
    left, right := 1, num
    for left <= right {
        mid := left + (right - left)/2
        if mid > num/mid {
            right = mid - 1
        } else {
            left = mid + 1
        }
    }
    if (left-1) * (left-1) == num {
        return true
    }
    return false
}