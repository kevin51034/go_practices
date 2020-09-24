// Time complexity: O(nlogn + n^2)
// Space complexity: O(1)

func threeSum(nums []int) [][]int {
    result := [][]int{}
    if len(nums)<3 {
        return result
    }
    sort.Ints(nums)
    for i:=0; i<len(nums); i++ {
		// prevent duplicate
        if i>0 && nums[i] == nums[i-1] {
            continue
        }
        target, left, right := -nums[i], i+1, len(nums)-1
        for left < right {
            sum := nums[left] + nums[right]
            if sum == target {
                result = append(result, []int{nums[i], nums[left], nums[right]})
                left++
				right--
				// prevent duplicate
                for left<right && nums[left] == nums[left-1] {
                    left++
                }
                for left<right && nums[right] == nums[right+1] {
                    right--
                }
            } else if sum < target {
                left++
            } else {
                right--
            }
        }
    }
    return result
}