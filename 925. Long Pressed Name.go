// Time complexity: O(n)
// Space complexity: O(1)

func isLongPressedName(name string, typed string) bool {
    left, right := 0, 0
    if len(name) == 0 {
        return true
    } else if len(typed) == 0 {
        return false
    }
    for left < len(name) && right < len(typed) {
        if name[left] != typed[right] {
            return false
        }
        for left<len(name)&&right<len(typed) && name[left]==typed[right] {
            left++
            right++
        }
        for right>0 && right < len(typed) && typed[right]==typed[right-1] {
            right++
        }
    }
    return left==len(name) && right==len(typed)
}