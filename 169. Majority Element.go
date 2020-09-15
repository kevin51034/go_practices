// Time compleixty: O(n)
// Space complexity: O(n)

func majorityElement(nums []int) int {
    numsMap := make(map[int]int, 0)
    for _,v := range nums {
        numsMap[v]++
        if numsMap[v] > len(nums)/2 {
            return v
        }
    }
    return 0
}

// Time compleixty: O(n)
// Space complexity: O(1)

func majorityElement(nums []int) int {
    result := nums[0]
    count := 0
    for _,v := range nums {
        if count == 0 {
            result = v
            count = 1
        } else {
            if result == v {
                count++
            } else {
                count--
            }
        }
    }
    return result
}