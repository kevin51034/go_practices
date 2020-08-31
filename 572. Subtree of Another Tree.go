/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSubtree(s *TreeNode, t *TreeNode) bool {
    if isSametree(s, t) {
        return true
    }
    if s == nil {
        return false
    }
    if isSubtree(s.Left, t) || isSubtree(s.Right, t) {
        return true
    }
    return false
}

func isSametree(s, t *TreeNode) bool {
    if s == nil && t == nil {
        return true
    }
    if s == nil || t == nil || s.Val!=t.Val {
        return false
    }
    return isSametree(s.Left, t.Left) && isSametree(s.Right, t.Right)
}