/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

 func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    if root == nil || root.Val == p.Val || root.Val == q.Val {
        return root
    }
    if p.Val > q.Val {
        p, q = q, p
    }
    if root.Val < p.Val {
        return lowestCommonAncestor(root.Right, p, q)
    } else if root.Val > q.Val {
        return lowestCommonAncestor(root.Left, p, q)
    }
    return root
}