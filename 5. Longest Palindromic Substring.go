// Time complexity: O(n^2)
// Space complexity: O(1)
// implement the concept of DP but not using extra space

func longestPalindrome(s string) string {
    result := ""
    for i:=0; i<len(s); i++ {
        isPalindrome(i, i, s, &result)
        isPalindrome(i, i+1, s, &result)
    }
    return result
}

func isPalindrome(l, r int, s string, result *string) {
    for l>=0 && r<len(s) {
        if s[l] != s[r] {
            return
        }
        if r-l+1 > len(*result) {
            *result = s[l:r+1]
        }
        l--
        r++
    }
    return
}