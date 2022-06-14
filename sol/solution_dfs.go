package sol

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraphDFS(node *Node) *Node {
	if node == nil {
		return nil
	}
	oldToNew := make(map[*Node]*Node)
	var dfs func(node *Node) *Node
	dfs = func(node *Node) *Node {
		if found, ok := oldToNew[node]; ok {
			return found
		}
		newNode := &Node{Val: node.Val}
		oldToNew[node] = newNode
		for _, nei := range node.Neighbors {
			newNode.Neighbors = append(newNode.Neighbors, dfs(nei))
		}
		return newNode
	}
	return dfs(node)
}
