/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

// solution 1
// iterative
func postorder(root *Node) []int {
    result := make([]int, 0)
    if root == nil {
        return result
    }
    stack := []*Node{root}
    for len(stack)>0 {
        node := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        result = append(result, node.Val)
        for i:=0; i<len(node.Children); i++ {
            stack = append(stack, node.Children[i])
        }
    }
    return reverse(result)
}

func reverse(result []int) []int {
    for i:=0; i<len(result)/2; i++ {
        result[i], result[len(result)-i-1] = result[len(result)-i-1], result[i]
    }
    return result
}

// solution 2
// recursive
func postorder(root *Node) []int {
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
    for _,node := range root.Children {
        travel(node, result)
    }
    *result = append(*result, root.Val)
    return
}