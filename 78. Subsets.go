// Time complexity : O(n * 2^n).
// n for every subset copy them into output list. and it has 2^n subset
// Space complexity :  O(n * 2^n).
// keep all the subsets of length n,
// since each of n elements could be present or absent.

func subsets(nums []int) [][]int {
    tempAns, result := []int{}, [][]int{}
    for sublen:=0; sublen<=len(nums);sublen++{
        backtracking(nums, 0, sublen, tempAns, &result)
    }
    return result
}

func backtracking(nums []int, start int, sublen int, tempAns []int, result *[][]int) {
    if len(tempAns) == sublen{
		ans := make([]int, sublen)
		// copy cost O(n) time
        copy(ans, tempAns)
        *result = append(*result, ans)
        return
    }
    
    for i:=start; i<(len(nums)-sublen+len(tempAns)+1); i++ {
        tempAns = append(tempAns, nums[i])
        backtracking(nums, i+1, sublen, tempAns, result)
        tempAns = tempAns[:len(tempAns)-1]
    }
    return
}