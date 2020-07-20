/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// reverse from levelorder

func levelOrderBottom(root *TreeNode) [][]int {
    tmp := levelOrder(root)
    result := make([][]int, 0)    
    for i:=len(tmp)-1; i>=0; i-- {
        result = append(result, tmp[i])
    }
    return result
}


// no.102 levelorder
func levelOrder(root *TreeNode) [][]int {
    result := make([][]int,0)
    queue := make([] *TreeNode,0)
    
    if root == nil {
        return result
    }
    queue = append(queue, root)
    
    for len(queue)>0 {
        list := make([]int, 0)
        l := len(queue)

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
        result = append(result, list)
    }
    return result
    
}