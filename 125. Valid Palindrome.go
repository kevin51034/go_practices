func isPalindrome(s string) bool {
    if len(s) == 0 {
        return true
    }
    s = strings.ToLower(s)
    left, right := 0, len(s)-1
    for left < right {
        for left<right && !(s[left]>='a' && s[left]<='z') && !(s[left]>='0' && s[left]<='9') {
            left++
        }
        for left<right && !(s[right]>='a' && s[right]<='z') && !(s[right]>='0' && s[right]<='9') {
            right--
        }
        if s[left] != s[right] {
            return false
        }
        left++
        right--
    }
    return true
}