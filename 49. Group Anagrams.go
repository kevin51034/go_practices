// solution without using sort
// time complexity: O(nk), n is len(strs), k is the length of the longest string in strs
// space complexity: O(nk), for we are using a hashmap

func groupAnagrams(strs []string) [][]string {
    result := make([][]string, 0)
    if len(strs) == 0 {
        return result
    }
    m := map[[26]int][]string{}
    for _, s := range strs {
        k := [26]int{}
        for i:=0; i<len(s); i++ {
            k[s[i] - 'a'] ++
        }
        m[k] = append(m[k], s)
    }
    for _,v := range m {
        result = append(result, v)
    }
    return result
}