// Time complexity: O(n)
// Space complexity: O(1)

func missingNumber(nums []int) int {
	num := 0
	for i := 0; i <= len(nums); i++ {
		num = num ^ i
	}

	for _, n := range nums {
		num = num ^ n
	}
	return num
}