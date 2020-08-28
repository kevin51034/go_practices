func isHappy(n int) bool {
    cycle := make(map[int]bool)
    return check(n, cycle)
}

func check(n int, cycle map[int]bool) bool {
    if n == 1 {
        return true
    }
    if cycle[n] == true {
        return false
    } else {
        cycle[n] = true
    }
    res := 0
    for n != 0 {
        tmp := n%10
        n /=10
        res += tmp*tmp
    }
    return check(res, cycle)
}