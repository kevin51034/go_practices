func thirdMax(nums []int) int {
    min := math.MinInt64
    one, two, three := min, min, min
    for i:=0; i<len(nums); i++ {
        if nums[i] > one {
            three = two
            two = one
            one = nums[i]
        } else if nums[i] < one && nums[i] > two {
            three = two
            two = nums[i]
        } else if nums[i] < two && nums[i] > three {
            three = nums[i]
        }
    }
    if three != min {
        return three
    } else {
        return one
    }
}