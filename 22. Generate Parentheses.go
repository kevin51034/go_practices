// Time complexity: O(2^n)
// Space complexity: O(k + n)

func generateParenthesis(n int) []string {
    result := make([]string, 0)
    dfs(n, n, "", &result)
    return result
}

func dfs(lindex, rindex int, str string, result *[]string) {
    if lindex==0 && rindex==0 {
        *result = append(*result, str)
        return
    }
    if lindex>0 {
        dfs(lindex-1, rindex, str+"(", result)
    }
    if rindex>0 && lindex<rindex {
        dfs(lindex, rindex-1, str+")", result)
    }
    return
}