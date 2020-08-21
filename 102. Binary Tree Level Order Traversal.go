/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// solution1
// using BFS
// Time complexity: O(n)
// Space complexity: O(n)

func levelOrder(root *TreeNode) [][]int {
    result := make([][]int, 0)
    queue := make([] *TreeNode, 0)
    
    if root == nil {
        return result
    }
    queue = append(queue, root)
    
    for len(queue)>0 {
        list := make([]int,0)
        l:=len(queue)

        for i:=0; i<l; i++ {
            thisNode := queue[0]
            queue = queue[1:]
            if thisNode.Left != nil {
                queue = append(queue, thisNode.Left)
            }
            if thisNode.Right != nil {
                queue = append(queue, thisNode.Right)
            }
            list = append(list, thisNode.Val)
        }
        result = append(result, list)
    }
    return result
}

// solution2
// using DFS
// Time complexity: O(n)
// Space complexity: O(n)

func levelOrder(root *TreeNode) [][]int {
    result := make([][]int, 0)
    if root == nil {
        return result
    }
    dfslevel(root, 0, &result)
    return result
}

func dfslevel(root *TreeNode, level int, result *[][] int) {
    if root == nil {
        return
    }
    curlevel := level + 1
    if len(*result)<curlevel {
        *result = append(*result, []int{})
    }
    (*result)[level] = append((*result)[level], root.Val)
    dfslevel(root.Left, curlevel, result)
    dfslevel(root.Right, curlevel, result)
}