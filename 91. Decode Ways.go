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
    if v,ok := visited[s]; ok {
        return v
    }
    if s[0] == '0' {
        return 0
    } else if len(s) == 1 {
        return 1
    }

    curway := findways(s[1:], visited)
    prefix,_ := strconv.Atoi(s[0:2])

    if prefix >= 10 &&  prefix <= 26 {
        curway += findways(s[2:], visited)
    }
    visited[s] = curway
    return curway
}

// solution 2
// DP
// Time complexity: O(n)
// Space complexity: O(1)

func numDecodings(s string) int {
    if len(s) == 0 || s[0] == '0' {
        return 0
    }
    if len(s) == 1 {
        return 1
    }
    one, two := 1, 1
    for i:=1; i<len(s); i++ {
        if s[i] == '0' {
            one = 0
        }
        num,_ := strconv.Atoi(s[i-1:i+1])
        if num >=10 && num <= 26 {
            tmp := one
            one += two
            two = tmp
        } else {
            two = one
        }
    }
    return one
}