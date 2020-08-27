func nextGreatestLetter(letters []byte, target byte) byte {
    left, right := 0, len(letters)-1
    for left <= right {
        mid := left + (right - left)/2
        if letters[mid] > target {
            right = mid - 1
        } else {
            left = mid + 1
        }
    }
    if left == len(letters) {
        return letters[0]
    }
    return letters[left]
}