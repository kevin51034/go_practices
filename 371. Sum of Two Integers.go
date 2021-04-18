// Time complexity: O(n)
// Space complexity: O(1), recursive will be slower because call stack

func getSum(a int, b int) int {
	// a is carry
	// b is remainder
	if a == 0 {
		return b
	}
	if b == 0 {
		return b
	}
	return getSum((a&b)<<1, a^b)
}

// After understanding that a ^ b stands for the sum of each digit (ignoring carry over)
// and a & b stands for whether carry over occurs on each digit,
// you will easily see that you need to add carry over (a & b) << 1
// repeatedly until no carry over occurred.
// (Because first carry over may produce another carry over,
// we should do this until no more carry over occurs)

// Time complexity: O(n)
// Space complexity: O(1)
func getSum(a int, b int) int {
	sum, carry := a^b, a&b
	for carry != 0 {
		carry = carry << 1
		sum, carry = sum^carry, sum&carry
	}
	return sum
}