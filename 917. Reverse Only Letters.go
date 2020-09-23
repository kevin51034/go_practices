func reverseOnlyLetters(S string) string {
    if len(S) == 0 {
        return S
    }
    s := []byte(S)
    left, right := 0, len(S)-1
    for left < right {
        for left<right && !(S[left]>='a' && S[left]<='z') && !(S[left]>='A' && S[left]<='Z') {
            left++
        }
        for left<right && !(S[right]>='a'&& S[right]<='z') && !(S[right]>='A' && S[right]<='Z') {
            right--
        }
        s[left], s[right] = s[right], s[left]
        left++
        right--
    }
    return string(s)
}