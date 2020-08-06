func coinChange(coins []int, amount int) int {
    f := make([]int, amount+1)
    f[0] = 0
    
    for i:=1; i<=amount; i++ {
        f[i] = amount+1
        for j:=0; j<len(coins); j++ {
            if coins[j]<=i {
                f[i] = min(f[i], f[i-coins[j]]+1)
            }
        }
    }
    if f[amount] == amount+1 {
        return -1
    }
    return f[amount]
}

func min(x, y int) int {
    if x<y {
        return x
    }
    return y
}