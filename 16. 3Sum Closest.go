// Time complexity: O(n^2)
// Space complexity: O(1)

func threeSumClosest(nums []int, target int) int {
    minDiff := math.MaxInt64
    result := 0
    sort.Ints(nums)
    for i:=0; i<len(nums)-2;i++ {
        if i>0 && nums[i] == nums[i-1] {
            continue
        }
        left, right := i+1, len(nums)-1
        for left<right {
            sum := nums[left] + nums[right] + nums[i]
            if abs(target, sum) < minDiff {
                minDiff = abs(target, sum)
                result = sum
            }
            if sum < target {
                left++
            } else {
                right--
            }
        }
    }
    return result
}

func abs(x, y int) int {
    if x>y {
        return x-y
    }
    return y-x
}