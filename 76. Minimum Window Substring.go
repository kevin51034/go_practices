// Time complexity: O(s + t)
// s for traversal, t for building hashmap
// Space complexity: O(s + t)

func minWindow(s string, t string) string {
    target := make(map[byte]int)
    window := make(map[byte]int)
    for i:=0; i<len(t); i++ {
        target[t[i]]++
    }
    
    left := 0
    right := 0
    match := 0
    ansLeft := 0
    ansRight := 0
    min := math.MaxInt64
    for right < len(s) {
        c := s[right]
        right++
        if target[c]!=0 {
            window[c]++
            if window[c] == target[c] {
                match++
            }
        }
        for match == len(target) {
            if right-left < min {
                min = right - left
                ansLeft = left
                ansRight = right
            }
            c = s[left]
            left++
            
            if target[c]!=0 {
                if window[c] == target[c] {
                    match--
                }
                window[c]--
            }
        }
    }
    return s[ansLeft:ansRight]
}