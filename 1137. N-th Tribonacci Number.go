// Time complexity: O(n)
// Space complexity: O(1)

func tribonacci(n int) int {
    if n == 0 {
        return 0
    } else if n == 1 || n == 2 {
        return 1
    }
    first, second, third := 0, 1, 1
    for i:=3; i<n; i++ {
        tmp := first + second + third
        first = second
        second = third
        third = tmp
    }
    return first + second + third
}