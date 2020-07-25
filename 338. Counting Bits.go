// Time complexity: O(n)
// Space complexity: O(n)

func countBits(num int) []int {
    bits := make([]int,num+1)
    
    for i:=1; i<len(bits); i++ {
        bits[i] = bits[i&(i-1)] +1
    }
    return bits
}