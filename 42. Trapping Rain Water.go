// DP solution
// Time complexity: O(n)
// Space complexity: O(n)

func trap(height []int) int {
    if len(height) == 0 {
        return 0
    }
    result := 0
    l := make([]int, len(height))
    r := make([]int, len(height))
    for i:=0; i<len(height); i++ {
        if i == 0 {
            l[i] = height[i]
            continue
        }
        l[i] = max(l[i-1], height[i])
    }
    for i:=len(height)-1; i>=0; i-- {
        if i == len(height)-1 {
            r[i] = height[i]
            continue
        }
        r[i] = max(r[i+1], height[i])
    }
    for i:=0; i<len(height); i++ {
        result += min(l[i], r[i]) - height[i]
    }
    return result
}

func max(x, y int) int {
    if x>y {
        return x
    }
    return y
}

func min(x, y int) int {
    if x<y {
        return x
    }
    return y
}

// Two pointer solution
// Time complexity: O(n)
// Space complexity: O(1)

func trap(height []int) int {
    if len(height) == 0 {
        return 0
    }
    maxLeft := height[0]
    maxRight := height[len(height)-1]
    left, right := 0, len(height)-1
    result := 0
    for left < right {
        if maxLeft < maxRight {
            result += maxLeft - height[left]
            left++
            maxLeft = max(maxLeft, height[left])
        } else {
            result += maxRight - height[right]
            right--
            maxRight = max(maxRight, height[right])
        }
    }
    return result
}

func max(x, y int) int {
    if x>y {
        return x
    }
    return y
}