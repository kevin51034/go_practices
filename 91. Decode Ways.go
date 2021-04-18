import "strconv"

// solution 1
// recursion with memoization
// Time complexity: O(n)
// Space complexity: O(n)

func numDecodings(s string) int {
	visited := make(map[string]int, 0)
	visited[""] = 1
	return findways(s, visited)
}

func findways(s string, visited map[string]int) int {
	if v, ok := visited[s]; ok {
		return v
	}
	if s[0] == '0' {
		return 0
	} else if len(s) == 1 {
		return 1
	}

	curway := findways(s[1:], visited)
	num, _ := strconv.Atoi(s[0:2])

	if num >= 10 && num <= 26 {
		curway += findways(s[2:], visited)
	}
	visited[s] = curway
	return curway
}

// solution 2
// DP
// Time complexity: O(n)
// Space complexity: O(1)
// implement the concept of fibonacci sequence
// this method come from dp
// dpone = dp[i-1], dptwo = dp[i-2]

/*
         First, we define dp[i] is the number of decode ways for input string s(0,1..,i-1,i) (i = 1,2 ..,len(s)-1).
		 Then we need to define our dp formula.
		 Examining the last 2 digits of s: s(i-1,i) = int(s[i-1]) * 10 + int(s[i])
         if 10 <= s(i-1,i) <= 26:
		                        ...i..  ...i..
             + s[i] == '0' (s = xx10xx, xx20xx, ...) --> dp[i] = dp[i-2]
             + s[i] != '0' (s = xx12xx, xx21xx, ...) --> dp[i] = dp[i-1] + dp[i-2]
         else if s(i-1,i) > 26 or s(i-1,i) < 10:
		                        ...i..  ...i..
             + s[i] == '0' (s = xx00xx, xx40xx..) --> invalid character --> return 0
             + s[i] != '0' (s = xx33xx, xx09xx..) -> dp[i] = dp[i-1]

          Because we just find dp[i], we can use two variables dptwo and dpone to store dp[i-2] and dp[i-1] respectively
*/

func numDecodings(s string) int {
	if len(s) == 0 || s[0] == '0' {
		return 0
	}
	dpone, dptwo := 1, 1
	for i := 1; i < len(s); i++ {
		if s[i] == '0' {
			dpone = 0
		}
		num, _ := strconv.Atoi(s[i-1 : i+1])
		if num >= 10 && num <= 26 {
			tmp := dpone
			dpone += dptwo
			dptwo = tmp
		} else {
			dptwo = dpone
		}
	}
	return dpone
}