// M doesn't have to pass in the pointer cuz it's slice

func findCircleNum(M [][]int) int {
    if len(M) == 0 {
        return 0
    }
    result := 0
    for i:=0; i<len(M); i++ {
        if M[i][i] == 0 {
            continue
        } else {
            result++
            dfs(i, M)
        }
    }
    return result
}

func dfs(i int, M [][] int) {
    for j:=0; j<len(M[i]); j++ {
        if M[i][j] == 0 {
            continue
        }
        M[i][j] = 0
        dfs(j, M)
    }
}

// with pointer, in case M is not slice but array

func findCircleNum(M [][]int) int {
    if len(M) == 0 {
        return 0
    }
    result := 0
    for i:=0; i<len(M); i++ {
        if M[i][i] == 0 {
            continue
        } else {
            result++
            dfs(i, &M)
        }
    }
    return result
}

func dfs(i int, M *[][]int) {
    if i<0 || i>=len((*M)) {
        return
    }
    for k:=0; k<len((*M)[i]); k++ {
        if (*M)[i][k] == 1 {
            (*M)[i][k] = 0
            dfs(k, M)
        }
    }
}
