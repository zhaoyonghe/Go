type UF struct {
    roots []int
    heights []int
}

func getUF(size int) *UF {
    a := make([]int, size)
    for i := 0; i < size; i++ {
        a[i] = i
    }
    return &UF {
        a,
        make([]int, size),
    }
}

func (uf *UF) root(i int) int {
    for uf.roots[i] != i {
        i = uf.roots[i]
    }
    return i
} 

func (uf *UF) union(a, b int) bool {
    ar := uf.root(a)
    br := uf.root(b)

    if ar == br {
        return false
    }

    if uf.heights[ar] > uf.heights[br] {
        uf.roots[br] = ar
    } else if uf.heights[ar] < uf.heights[br] {
        uf.roots[ar] = br
    } else {
        uf.roots[ar] = br
        uf.heights[br]++
    }

    return true
}

func findRedundantDirectedConnection(edges [][]int) []int {
    n := len(edges)

    in := make([]int, n + 1)
    for i := 0; i <= n; i++ {
        in[i] = -1
    }
    
    // check the indegrees
    for i, edge := range edges {
        if in[edge[1]] == -1 {
            in[edge[1]] = i
        } else {
            temp := append(edges[:i], edges[i+1:]...)
            if has, _ := hasCycle(temp); has {
                return edges[in[edge[1]]]
            } else {
                return edge
            }
        }
    }

    // no node has 2 indegrees
    _, res := hasCycle(edges)
    return res
}

func hasCycle(edges [][]int) (bool, []int) {
    n := len(edges)
    uf := getUF(n + 2)

    for _, edge := range edges {
        if !uf.union(edge[0], edge[1]) {
            return true, edge
        }
    }

    return false, []int{}
}