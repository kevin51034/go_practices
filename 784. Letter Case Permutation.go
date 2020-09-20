// DFS
// Time complexity: O(n*2^l), l = # of letters in the string
// Space complexity: O(n) + O(n*2^l)

func letterCasePermutation(S string) []string {
    result := make([]string, 0)
    dfs(S, "", 0, &result)
    return result
}

func dfs(S string, set string, index int, result *[]string) {
    if index == len(S) {
        *result = append(*result, set)
        return
    }
    if S[index] >= 48 && S[index] <= 57 {
        dfs(S, set+string(S[index]), index+1, result)
    } else if S[index] >= 97 {
        dfs(S, set+string(S[index]), index+1, result)
        dfs(S, set+string(S[index]-32), index+1, result)
    } else {
        dfs(S, set+string(S[index]), index+1, result)
        dfs(S, set+string(S[index]+32), index+1, result)
    }
    return
}