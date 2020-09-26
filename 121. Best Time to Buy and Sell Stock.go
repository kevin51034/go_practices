// Time complexity: O(n)
// Space complexity: O(1)

func maxProfit(prices []int) int {
    if len(prices) <= 1 {
        return 0
    }
    result := 0
    min := prices[0]
    for i:=0; i<len(prices); i++ {
        if prices[i] - min > result {
            result = prices[i] - min
        }
        if prices[i] < min {
            min = prices[i]
        }
    }
    return result
}