// Time complexity : 
// Assume that n = length of haystack and m = length of needle,
// then the runtime complexity is O(nm).
// Space complexity :  O(1)

func strStr(haystack string, needle string) int {
    if len(needle) == 0 {
        return 0
    }
    for i:=0; i<(len(haystack) - len(needle)) +1; i++ {
		j:=0
		for ; j<len(needle); j++ {
			if haystack[i+j] != needle[j] {
				break;
			}
		}
		if j == len(needle){
			return i
		}
	}
	return -1
}