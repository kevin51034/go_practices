func isIsomorphic(s string, t string) bool {
    s1, s2 := [256]int{}, [256]int{}
    for i:=0; i<len(s); i++ {
        if s1[s[i]]!=s2[t[i]] {
            return false
        }
        s1[s[i]] = i+1
        s2[t[i]] = i+1
    }
    return true
}