/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// Time complexity: O(n)
// Space complexity: O(n)

func zigzagLevelOrder(root *TreeNode) [][]int {
    result := make([][]int, 0)
    if root == nil {
        return result
    }
    queue := make([] *TreeNode, 0)
    
    queue = append(queue, root)
    toggle := false
    for len(queue)>0 {

        l:= len(queue)
        list := make([]int,0)
        
        for i:=0; i<l; i++ {
            thisNode := queue[0]
            queue = queue[1:]
            list = append(list, thisNode.Val)
            if thisNode.Left != nil {
                queue = append(queue, thisNode.Left)
            }
            if thisNode.Right != nil {
                queue = append(queue, thisNode.Right)
            }
        }
        if toggle {
            reverse(list)
        }
        result = append(result, list)
        toggle = !toggle
    }
    return result
    
}

func reverse(list []int) {
    for i:=0; i<len(list)/2; i++ {
        list[i],list[len(list)-i-1] = list[len(list)-i-1], list[i]
    }
}
