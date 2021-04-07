// Time complexity: O(n)
// Space complexity: O(1)

func maxProfit(prices []int) int {
    preMin := prices[0]
    result := 0
    for i:=1; i<len(prices); i++ {
        if prices[i] < preMin {
            preMin = prices[i]
        } else {
            result = max(result, prices[i]-preMin)
        }
    }
    return result
}

func max(x, y int) int {
    if x>y {
        return x
    }
    return y
}