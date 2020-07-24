/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

// Time complexity: O(n)
// Space complexity: O(n)

func cloneGraph(node *Node) *Node {
    visited := make(map[*Node]*Node, 0)
    return clone(node, visited)
}

func clone(node *Node, visited map[*Node]*Node) *Node {
    if node == nil {
        return node
    }
    
    if val,ok:=visited[node];ok {
        return val
    }
    
    newNode := &Node{Val: node.Val,Neighbors: make([]*Node,len(node.Neighbors))}
    visited[node] = newNode
    for i:=0;i<len(node.Neighbors);i++ {
        newNode.Neighbors[i] = clone(node.Neighbors[i], visited)
    }
    return newNode
}