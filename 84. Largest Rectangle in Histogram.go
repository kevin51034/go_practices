func largestRectangleArea(heights []int) int {
    if len(heights) == 0 {
        return 0
    }
    height, maxArea := 0, 0
    stack := make([]int, 0)
    for i:=0; i<=len(heights); i++ {
        if i==len(heights) {
            height = 0
        } else {
            height = heights[i]
        }
        if len(stack)==0 || height>heights[stack[len(stack)-1]] {
            stack = append(stack, i)
        } else {
            tmp:=stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            
            width:=0
            if len(stack) == 0 {
                width = i
            } else {
                width = i - 1 - stack[len(stack)-1]
            }
			maxArea = max(maxArea, width*heights[tmp])
			// need to reduce i to redo index i 
			// because above is to calculate i-1 area
			// need to push height back into stack
            i--
        }
    }
    return maxArea
}

func max (x int, y int) int {
    if x>y {
        return x
    }
    return y
}