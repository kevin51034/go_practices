// Time complexity: O(n)
// Space complexity: O(1)

func maxArea(height []int) int {
    result := 0
    if len(height) == 0 {
        return 0
    }
    left, right := 0, len(height)-1
    for left < right {
        minHeight := min(height[left], height[right])
        result = max(minHeight*(right-left), result)
        if height[left] < height[right] {
            left++
        } else {
            right--
        }
    }
    return result
}

func min(x, y int) int {
    if x<y {
        return x
    }
    return y
}

func max(x, y int) int {
    if x>y {
        return x
    }
    return y
}