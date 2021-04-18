// Time complexity: O(n)
// Space complexity: O(n)

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sm := [26]int{}
	tm := [26]int{}
	for i := 0; i < len(s); i++ {
		sm[s[i]-'a']++
		tm[t[i]-'a']++
	}
	return sm == tm
}