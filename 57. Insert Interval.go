// solution 1
// Time complexity: O(n)
// Space complexity: O(n)
// keep updating the newInterval

func insert(intervals [][]int, newInterval []int) [][]int {
    result := make([][]int, 0)
    i := 0
    for i<len(intervals) && intervals[i][1] < newInterval[0] {
        result = append(result, intervals[i])
        i++
        continue
    }
    // overlapping
    for i<len(intervals) && newInterval[1] >= intervals[i][0] {
        newInterval[0] = min(newInterval[0], intervals[i][0])
        newInterval[1] = max(newInterval[1], intervals[i][1])
        i++
    }
    result = append(result, newInterval)
    for i<len(intervals) {
        result = append(result, intervals[i])
        i++
    }
    return result
}

func min(x, y int) int {
    if x<y {
        return x
    }
    return y
}

func max(x, y int) int {
    if x>y {
        return x
    }
    return y
}


// solution 2
// Time complexity: O(n)
// Space complexity: O(n)
// extend from 56. Merge Intervals
// first insert the newInterval than merge them
// but it's quite hard to insert the newInterval using slice so just use solution 1

func insert(intervals [][]int, newInterval []int) [][]int {
    result := make([][]int, 0)
    if len(intervals) == 0 {
        return append(result, newInterval)
    }
	// this section can code better
    for i:=len(intervals)-1; i>=0; i-- {
        if intervals[i][0] <= newInterval[0] {
            tmp := append(make([][]int, 0), intervals[i+1:]...)
            intervals = append(intervals[:i+1], newInterval)
            intervals = append(intervals, tmp...)
            break
        }
        if i==0 {
            intervals = append(intervals, newInterval)
            copy(intervals[1:], intervals)
            intervals[0] = newInterval
        }
    }
    for _,inter := range intervals {
        if len(result) == 0 {
            result = append(result, inter)
            continue
        }
        preInter := result[len(result)-1]
        if preInter[1] < inter[0] {
            result = append(result, inter)
        } else if preInter[1] >= inter[0] {
            result[len(result)-1][1] = max(preInter[1], inter[1])
        }
    }
        return result
}

func max(x, y int) int {
    if x>y {
        return x
    }
    return y
}