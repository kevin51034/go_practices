func findDuplicate(nums []int) int {
    fast := nums[nums[0]]
    slow := nums[0]
    for fast != slow {
        fast = nums[nums[fast]]
        slow = nums[slow]
    }
    fast = 0
    for slow != fast {
        fast = nums[fast]
        slow = nums[slow]
    }
    return slow
}