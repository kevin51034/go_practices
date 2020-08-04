// Time complexity: O(mn), m n = len(word1) len(word2)
// Space complexity: O(mn)

func minDistance(word1 string, word2 string) int {
    f := make([][]int, len(word1)+1)
    for i:=0; i<len(f); i++ {
        f[i] = make([]int, len(word2)+1)
    }
    for i:=0; i<len(f); i++ {
        f[i][0] = i
    }
    for j:=0; j<len(f[0]); j++ {
        f[0][j] = j
    }
    for i:=1; i<len(f); i++ {
        for j:=1; j<len(f[i]); j++ {
            if word1[i-1] == word2[j-1] {
                f[i][j] = f[i-1][j-1]
            } else {
                f[i][j] = min(f[i-1][j-1], f[i-1][j], f[i][j-1]) + 1
            }
        }
    }
    for i:=0; i<=len(word1); i++ {
        fmt.Println(f[i])
    }
    return f[len(word1)][len(word2)]
}

func min(x, y, z int) int {
    if x<=y && x<=z {
        return x
    } else if y<=x && y<=z {
        return y
    }
    return z
}