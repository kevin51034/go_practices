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

// optimize
// Time complexity: O(mn), m n = len(text1) len(text2)
// Space complexity: O(n)

func longestCommonSubsequence(text1 string, text2 string) int {
    f := make([][]int, 2)
    
    for i:=0; i<=len(text1); i++ {
        f[0] = make([]int, len(text2)+1)
        for j:=0; j<=len(text2); j++ {
            if i==0 || j==0 {
                continue
            }
            if text1[i-1] == text2[j-1] {
                f[0][j] = f[1][j-1] + 1
            } else {
                f[0][j] = max(f[1][j], f[0][j-1])
            }
        }
        f[0], f[1] = f[1], f[0]
    }
    return f[1][len(text2)]
}

func max(x ,y int) int {
    if x>y {
        return x
    }
    return y
}