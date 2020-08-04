func wordBreak(s string, wordDict []string) bool {
    f := make([]bool, len(s)+1)
    for i,_ := range f {
        f[i] = false
    }
    f[0] = true
    
    for i:=1; i<=len(s); i++ {
        for j:=0; j<i; j++ {
            if f[j] && inDict(s[j:i], wordDict) {
                f[i] = true
                break
            }
        }
    }
    return f[len(s)]
}

func inDict(s string, wordDict []string) bool {
    for _,v := range wordDict {
        if s == v {
            return true
        }
    }
    return false
}