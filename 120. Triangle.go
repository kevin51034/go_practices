// solution 1
// using top-down dp with extra space O(n^2)
// Time complexity: O(n^2)
// Space complexity: O(n^2)

func minimumTotal(triangle [][]int) int {
    f := make([][]int, len(triangle))
    for i:=0; i<len(triangle); i++ {
        f[i] = make([]int, i+1)
        for j:=0; j<=i; j++ {
            f[i][j] = triangle[i][j]
            if i==0 && j==0 {
                continue
            }
            if j==0 {
                f[i][j] = f[i-1][j] + triangle[i][j]
            } else if j == i {
                f[i][j] = f[i-1][j-1] + triangle[i][j]
            } else {
                f[i][j] = min(f[i-1][j-1], f[i-1][j]) + triangle[i][j] 
            }
        }
    }
    return minElement(f[len(triangle)-1])
}

func min(x, y int) int {
    if x<y {
        return x
    }
    return y
}

func minElement(a []int) int {
    min := math.MaxInt64
    for _,v := range a {
        if v<min {
            min = v
        }
    }
    return min
}


// solution 2
// using top-down dp with extra space O(n)
// Time complexity: O(n^2)
// Space complexity: O(n^2) // because original triangle space is O(n^2)

func minimumTotal(triangle [][]int) int {
    f := make([][]int, 2)
    for i:=0; i<len(triangle); i++ {
        f[0] = make([]int, i+1)
        for j:=0; j<=i; j++ {
            f[0][j] = triangle[i][j]
            if i==0 && j==0 {
                continue
            }
            if j==0 {
                f[0][j] = f[1][j] + triangle[i][j]
            } else if j==i {
                f[0][j] = f[1][j-1] + triangle[i][j]
            } else {
                f[0][j] = min(f[1][j-1], f[1][j]) + triangle[i][j]
            }
        }
        f[0], f[1] = f[1], f[0]
    }
    return minElement(f[1])
}

func min(x, y int) int {
    if x<y {
        return x
    }
    return y
}

func minElement(a []int) int {
    min := math.MaxInt64
    for _,v := range a {
        if v<min {
            min = v
        }
    }
    return min
}


// solution 3
// using top-down dp with extra space O(1)
// Time complexity: O(n^2)
// Space complexity: O(n^2) // because original triangle space is O(n^2)

func minimumTotal(triangle [][]int) int {
    for i:=0; i<len(triangle); i++ {
        for j:=0; j<=i; j++ {
            if i==0 && j==0 {
                continue
            }
            if j==0 {
                triangle[i][j] = triangle[i-1][j] + triangle[i][j]
            } else if j==i {
                triangle[i][j] = triangle[i-1][j-1] + triangle[i][j]
            } else {
                triangle[i][j] = min(triangle[i-1][j-1], triangle[i-1][j]) + triangle[i][j]
            }
        }
    }
    return minElement(triangle[len(triangle)-1])
}

func min(x, y int) int {
    if x<y {
        return x
    }
    return y
}

func minElement(a []int) int {
    min := math.MaxInt64
    for _,v := range a {
        if v<min {
            min = v
        }
    }
    return min
}


// solution 4
// using bottom-up dp with extra space O(n^2)
// Time complexity: O(n^2)
// Space complexity: O(n^2)

func minimumTotal(triangle [][]int) int {
    f := make([][]int, len(triangle))
    for i:=len(triangle)-1; i>=0; i-- {
        f[i] = make([]int, i+1)
        for j:=0; j<=i; j++ {
            if i==len(triangle)-1 {
                f[i] = triangle[i]
                continue
            }
            f[i][j] = min(f[i+1][j], f[i+1][j+1]) + triangle[i][j] 
        }
    }
    return f[0][0]
}

func min(x, y int) int {
    if x<y {
        return x
    }
    return y
}


// solution 5
// using bottom-up dp with extra space O(n)
// Time complexity: O(n^2)
// Space complexity: O(n^2) // because original triangle space is O(n^2)

func minimumTotal(triangle [][]int) int {
    f := make([][]int, 2)
    for i:=len(triangle)-1; i>=0; i-- {
        f[0] = make([]int, i+1)
        for j:=0; j<=i; j++ {
            if i==len(triangle)-1 {
                f[0] = triangle[i]
                continue
            }
            f[0][j] = min(f[1][j], f[1][j+1]) + triangle[i][j] 
        }
        f[0], f[1] = f[1], f[0]
    }
    return f[1][0]
}

func min(x, y int) int {
    if x<y {
        return x
    }
    return y
}


// solution 6
// using bottom-up dp with extra space O(1)
// Time complexity: O(n^2)
// Space complexity: O(n^2) // because original triangle space is O(n^2)

func minimumTotal(triangle [][]int) int {
    for i:=len(triangle)-1; i>=0; i-- {
        for j:=0; j<=i; j++ {
            if i==len(triangle)-1 {
                continue
            }
            triangle[i][j] = min(triangle[i+1][j], triangle[i+1][j+1]) + triangle[i][j]
        }
    }
    return triangle[0][0]
}

func min(x, y int) int {
    if x<y {
        return x
    }
    return y
}