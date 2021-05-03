// Time compleixty: O(2^n)
// Space complexity: O(n)

func combinationSum(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	set := make([]int, 0)
	findCombinations(candidates, 0, target, set, &result)
	return result
}

func findCombinations(nums []int, index, target int, set []int, result *[][]int) {
	if target == 0 {
		tmp := make([]int, len(set))
		copy(tmp, set)
		*result = append(*result, tmp)
		return
	} else if target < 0 {
		return
	}
	for i := index; i < len(nums); i++ {
		set = append(set, nums[i])
		findCombinations(nums, i, target-nums[i], set, result)
		set = set[:len(set)-1]
	}
	return
}