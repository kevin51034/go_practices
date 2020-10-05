// time complexity: O(n)
// space complexity: O(1)

func sortColors(nums []int)  {
    red, blue := 0, 0
    for i:=0; i<len(nums)-blue; i++ {
        if nums[i] == 0 {
            nums[i], nums[red] = nums[red], nums[i]
            red++
        } else if nums[i] == 2 {
            nums[i], nums[len(nums)-blue-1] = nums[len(nums)-blue-1], nums[i]
            i--
            blue++
        }
    }
}