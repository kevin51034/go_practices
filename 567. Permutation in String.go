// Time complexity: O(s + t)
// Space complexity: O(1)
// O(1) because for the input two strings only contain lower case letters. O(26*2) = O(1)

func checkInclusion(s1 string, s2 string) bool {
    target := make(map[byte]int)
    window := make(map[byte]int)
    for i:=0; i<len(s1); i++ {
        target[s1[i]]++
    }
    
    left, right, match := 0, 0, 0
    
    for right < len(s2) {
        c := s2[right]
        right++
        if target[c] != 0 {
            window[c]++
            if window[c] == target[c] {
                match ++
            }
        }
        
        for right - left >= len(s1) {
            if match == len(target) {
                return true
            }
            
            d := s2[left]
            left ++
            if target[d] != 0 {
                if target[d] == window[d] {
                    match --
                }
                window[d] --
            }
        }
    }
    return false
}