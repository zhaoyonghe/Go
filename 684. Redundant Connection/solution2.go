type UF struct {
    roots []int
    orders []int
}

func newUF(n int) *UF {
    tmp := make([]int, n)
    for i := 0; i < n; i++ {
        tmp[i] = i
    }
    return &UF{
        roots: tmp,
        orders: make([]int, n),
    }
}

func (uf *UF) root(a int) int {
    for uf.roots[a] != a {
        a = uf.roots[a]
    }
    return a
}

func (uf *UF) union(a, b int) {
    i := uf.root(a)
    j := uf.root(b)
    if i == j {
        return
    }

    if uf.orders[i] < uf.orders[j] {
        uf.roots[i] = j
    } else {
        uf.roots[j] = i
        if uf.orders[i] == uf.orders[j] {
            uf.orders[i]++
        }
    }
}

func (uf *UF) same(a, b int) bool {
    return uf.root(a) == uf.root(b)
}

func findRedundantConnection(edges [][]int) []int {
    n := len(edges)
    uf := newUF(n + 1)

    for _, edge := range edges {
        if uf.same(edge[0], edge[1]) {
            return edge
        } else {
            uf.union(edge[0], edge[1])
        }
    }
    return []int{0, 0}
}