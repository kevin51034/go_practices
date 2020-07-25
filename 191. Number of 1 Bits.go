// Time complexity: O(1)
// Space complexity: O(1)

func hammingWeight(num uint32) int {
    count := 0
    for num!= 0 {
        num = num&(num-1)
        count++
    }
    return count
}