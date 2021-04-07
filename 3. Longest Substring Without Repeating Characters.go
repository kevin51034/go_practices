// Time complexity: O(n)
// Space complexity: O(1)

func lengthOfLongestSubstring(s string) int {
    window := make(map[byte]int)
    
    left, right, result := 0, 0, 0
    for right < len(s) {
        c := s[right]
        right++
        window[c]++
        for window[c]>1 {
            d := s[left]
            left++
            window[d] --
        }
        result = max(result, right-left)
    }
    return result
}

func max(x, y int) int {
    if x>y {
        return x
    }
    return y
}

// optimize

func lengthOfLongestSubstring(s string) int {
    index := [128]int{}
    ans := 0
    for i:=0; i<128; i++ {
        index[i] = -1
	}
	// j is new character
	// i is prev position of j character
    for i,j:=0,0; j<len(s); j++ {
        i = max(i, index[s[j]]+1)
        ans = max(ans, j-i+1)
        index[s[j]] = j
    }
    return ans
}

func max(x, y int) int {
    if x>y {
        return x
    }
    return y
}