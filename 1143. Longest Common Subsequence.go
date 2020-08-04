// Time complexity: O(mn), m n = len(text1) len(text2)
// Space complexity: O(mn)

func longestCommonSubsequence(text1 string, text2 string) int {
    f := make([][]int, len(text1)+1)
    for i:=0; i<=len(text1); i++ {
        f[i] = make([]int, len(text2)+1)
    }
    for i:=1; i<=len(text1); i++ {
        for j:=1; j<=len(text2); j++ {
            if text1[i-1] == text2[j-1] {
                f[i][j] = f[i-1][j-1] + 1
            } else {
                f[i][j] = max(f[i-1][j], f[i][j-1])
            }
        }
    }
    for i:=0; i<=len(text1); i++ {
        fmt.Println(f[i])
    }
    return f[len(text1)][len(text2)]
}

func max(x, y int) int {
    if x>y {
        return x
    }
    return y
}