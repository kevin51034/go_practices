// Time complexity: O(n) worse case like 222222222
// Space complexity: O(n)

func findMin(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    left := 0
    right := len(nums) - 1
    for left<right {
        for left < right && nums[right-1] == nums[right] {
            right--
        }
        for left < right && nums[left] == nums[left+1] {
            left ++
        }
        mid := (left+right)/2
        if nums[mid] > nums[right] {
            left = mid + 1
        } else {
            right = mid
        }
    }
    if nums[left] > nums[right] {
        return nums[right]
    }
    return nums[left]
}