// Time complexity: O(n)
// Space complexity: O(1)

func findAnagrams(s string, p string) []int {
    target := make(map[byte]int)
    window := make(map[byte]int)
    for i:=0; i<len(p); i++ {
        target[p[i]]++
    }
    left, right, match := 0, 0, 0
    result := make([]int, 0)
    for right < len(s) {
        c := s[right]
        right ++
        if target[c] != 0 {
            window[c]++
            if window[c] == target[c] {
                match ++
            }
        }
        fmt.Println(window)
        for right-left >= len(p) {
            if match == len(target) && right-left == len(p) {
                result = append(result, left)
            }
            
            d := s[left]
            left++
            
            if target[d] != 0 {
                if window[d] == target[d] {
                    match --
                }
                window[d] --
            }
        }
    }
    return result
}


// optimize
// Time complexity: O(n)
// Space complexity: O(1)

func findAnagrams(s string, p string) []int {
    target := [26]int{}
    window := [26]int{}
    
    for i:=0; i<len(p); i++ {
        target[p[i]-'a']++
    }
    result := make([]int, 0)
    for i,_ := range s {
        window[s[i]-'a'] ++
        if i >= len(p) {
            window[s[i-len(p)]-'a'] --
        }
        if window == target {
            result = append(result, i-len(p)+1)
        }
    }
    return result
}
