/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

 func levelOrder(root *Node) [][]int {
    result := make([][]int, 0)
    if root == nil {
        return result
    }
    queue := []*Node{root}
    for len(queue)>0 {
        length := len(queue)
        list := make([]int, 0)
        for i:=0; i<length; i++ {
            node := queue[0]
            queue = queue[1:]
            list = append(list, node.Val)
            if len(node.Children)>0 {
                for j:=0; j<len(node.Children); j++ {
                    queue = append(queue, node.Children[j])
                }
            }
        }
        result = append(result, list)
    }
    return result
}