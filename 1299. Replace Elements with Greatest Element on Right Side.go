func replaceElements(arr []int) []int {
    tmpmax := -1
    for i:=len(arr)-1; i>=0; i-- {
        if i == len(arr)-1 {
            tmpmax = arr[i]
            arr[i] = -1
            continue
        }
        tmp := max(tmpmax, arr[i])
        arr[i] = tmpmax
        tmpmax = tmp
    }
    return arr
}

func max(x, y int) int {
    if x>y {
        return x
    }
    return y
}