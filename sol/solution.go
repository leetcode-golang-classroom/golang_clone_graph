package sol

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	visit := make(map[int]*Node)
	created := make(map[int]*Node)
	var result *Node
	var bfs = func(node *Node) {
		queue := []*Node{node}
		for len(queue) != 0 {
			top := queue[0]
			queue = queue[1:]
			// mark as visited
			visit[top.Val] = top
			// create top if no existed
			var cur *Node
			if found, ok := created[top.Val]; ok {
				cur = found
			} else {
				cur = &Node{Val: top.Val}
				created[cur.Val] = cur
			}
			// assign to result if cur.Val = 1
			if cur.Val == 1 {
				result = cur
			}
			// copy Neighbors if exists
			nLen := len(top.Neighbors)
			if nLen != 0 {
				cur.Neighbors = make([]*Node, nLen)
				for idx := range top.Neighbors {
					n := top.Neighbors[idx]
					var newNode *Node
					if find, ok := created[n.Val]; ok {
						newNode = find
					} else {
						newNode = &Node{Val: n.Val}
						created[n.Val] = newNode
					}
					cur.Neighbors[idx] = newNode
					// check if not visited push to queue
					if _, ok := visit[n.Val]; !ok {
						queue = append(queue, n)
					}
				}
			}
		}
	}
	bfs(node)
	return result
}
