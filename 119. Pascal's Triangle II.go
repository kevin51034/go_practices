func getRow(rowIndex int) []int {
    f := make([][]int, 2)
    for i:=0; i<=rowIndex; i++ {
        f[0] = make([]int, i+1)
        for j:=0; j<=i; j++ {
            if j==0 || j==i {
                f[0][j] = 1
            } else {
                f[0][j] = f[1][j-1] + f[1][j]
            }
        }
        f[0], f[1] = f[1], f[0]
    }
    return f[1]
}