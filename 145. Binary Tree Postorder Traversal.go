/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// solution 1
// recursive

func postorderTraversal(root *TreeNode) []int {
    result := make([]int, 0)
    travel(root, &result)
    return result
}

func travel(root *TreeNode, result *[]int) {
    if root == nil {
        return
    }
    travel(root.Left, result)
    travel(root.Right, result)
    *result = append(*result, root.Val)
}

// solution 1
// iteration
// just like reverse preorder travesal

func postorderTraversal(root *TreeNode) []int {
    result := make([]int, 0)
    stack := make([]*TreeNode, 0)
    if root == nil {
        return result
    }
    stack = append(stack, root)
    for len(stack)>0 {
        node := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        if node != nil {
            result = append(result, node.Val)
        }
        if node.Left != nil {
            stack = append(stack, node.Left)
        }
        if node.Right != nil {
            stack = append(stack, node.Right)
        }
    }
    return reverse(result)
}

func reverse(result []int) [] int {
    for i:=0; i<len(result)/2; i++ {
        result[i], result[len(result)-i-1] = result[len(result)-i-1], result[i]
    }
    return result
}