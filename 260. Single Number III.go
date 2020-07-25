// using last set bit (lsb) to seperate to two group
// Time complexity: O(n)
// Space complexity: O(1)

func singleNumber(nums []int) []int {
    sum := 0
    
    for _,num := range nums {
        sum ^= num
    }
    sum &= -sum
    result := []int {0, 0}
    for _,num := range nums {
        if sum & num == 0 {
            result[0] ^= num
        } else {
            result[1] ^= num
        }
    }
    return result
}