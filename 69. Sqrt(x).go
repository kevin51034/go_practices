func mySqrt(x int) int {
    left, right := 1, x
    for left <= right {
        mid := left + (right - left)/2
        if mid * mid == x {
            return mid
        } else if mid * mid < x {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    if left * left > x {
        return left - 1
    }
    return left
}