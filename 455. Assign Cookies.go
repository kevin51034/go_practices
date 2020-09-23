func findContentChildren(g []int, s []int) int {
    if len(g)==0 || len(s)==0 {
        return 0
    }
    result := 0
    sort.Ints(g)
    sort.Ints(s)
    for i, j:=0, 0; i<len(g)&&j<len(s);{
        if g[i]<=s[j] {
            result++
            i++
            j++
        } else {
            j++
        }
    }
    return result
}