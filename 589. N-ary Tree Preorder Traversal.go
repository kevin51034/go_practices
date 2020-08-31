/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

// solution 1
// recursive
func preorder(root *Node) []int {
    result := make([]int, 0)
    if root == nil {
        return result
    }
    travel(root, &result)
    return result
}

func travel(root *Node, result *[]int) {
    if root == nil {
        return
    }
    *result = append(*result, root.Val)
    for _,node := range root.Children {
        travel(node, result)
    }
}

// solution 2
// iterative
func preorder(root *Node) []int {
    result := make([]int, 0)
    if root == nil {
        return result
    }
    stack := []*Node{root}
    for len(stack)>0 {
        node := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        result = append(result, node.Val)
        for i:=len(node.Children)-1; i>=0; i-- {
            stack = append(stack, node.Children[i])
        }
    }
    return result
}