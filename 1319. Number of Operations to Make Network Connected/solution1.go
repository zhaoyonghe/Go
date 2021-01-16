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

func makeConnected(n int, connections [][]int) int {
    c := len(connections)
    if c < n - 1 {
        return -1
    }

    uf := getUF(n)
    for _, edge := range connections {
        uf.union(edge[0], edge[1])
    }

    m := make(map[int]bool)
    for i := 0; i < n; i++ {
        m[uf.root(i)] = true
    }

    return len(m) - 1
}

