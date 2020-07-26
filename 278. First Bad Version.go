/** 
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad 
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
    left := 1
    right := n
    if n <= 1 {
        return n
    }
    mid := 1
    for left<=right {
        mid := (left+right)/2
        if isBadVersion(mid) && !isBadVersion(mid-1) {
            return mid
        } else if isBadVersion(mid) {
            right = mid - 1
        } else {
            left = mid + 1
        }
    }
    return mid
}

// solution2
// reduce the API calls

func firstBadVersion(n int) int {
    left := 1
    right := n
    if n <= 1 {
        return n
    }
    for left < right {
        mid := (left+right)/2
        if isBadVersion(mid){
            right = mid
        } else {
            left = mid + 1
        }
    }
    return left
}