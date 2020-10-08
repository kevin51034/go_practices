// time complexity: O(nlogn)
// space complexity: O(1)

func merge(intervals [][]int) [][]int {
    result := make([][]int, 0)
    if len(intervals) == 0 {
        return result
    }
    sort.Slice(intervals, func(i,j int)bool{return intervals[i][0]<intervals[j][0]})
    for i,inter := range intervals {
        if i==0 {
            result = append(result, inter)
            continue
        }
        preInter := result[len(result)-1]
        if inter[0] <= preInter[1] && inter[1] > preInter[1] {
            result[len(result)-1][1] = inter[1]
        } else if inter[0] > preInter[1] {
            result = append(result, inter)
        }
    }
    return result
}