type Node struct {
    Val int
    Neighbors []*Node
}


func cloneGraph(node *Node) *Node {
    if node == nil {
        return nil
    }

    var status map[*Node]int = make(map[*Node]int)
    var m map[*Node]*Node = make(map[*Node]*Node)

    m[node] = &Node{Val:node.Val,Neighbors:[]*Node{}}

    dfs(node, status, m)

    return m[node]
}


func dfs(start *Node, status map[*Node]int, m map[*Node]*Node) {
    status[start] = 1
    var cstart *Node = m[start]

    for _, nei := range start.Neighbors {
        if cnei, ok := m[nei]; ok {
            cstart.Neighbors = append(cstart.Neighbors, cnei)
        } else {
            var tempcnei *Node = &Node{Val:nei.Val,Neighbors:[]*Node{}}
            cstart.Neighbors = append(cstart.Neighbors, tempcnei)
            m[nei] = tempcnei
        }
        if _, ok := status[nei]; !ok {
            dfs(nei, status, m)
        }        
    }
}