func checkIfExist(arr []int) bool {
    list := make(map[int]bool)
    for _,v := range arr {
        if v%2 == 0 && list[v/2] {
            return true
        }
        if list[2*v] {
            return true
        }
        list[v] = true
    }
    return false
}