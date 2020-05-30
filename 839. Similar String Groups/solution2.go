type UF struct {
    parents []int
    heights []int
    count int
}

func newUF(n int) *UF {
    a := make([]int, n)
    for i := 0; i < n; i++ {
        a[i] = i
    }
    b := make([]int, n)
    return &UF{
        a,
        b,
        n,
    }
}

func (uf *UF) root(a int) int {
    for {
        p := uf.parents[a]
        if p == a {
            return p
        }
        a = p
    }
}

func (uf *UF) union(a, b int) {
    ra := uf.root(a)
    rb := uf.root(b)

    if ra == rb {
        return
    }

    ha := uf.heights[ra]
    hb := uf.heights[rb]

    if ha > hb {
        uf.parents[rb] = ra
    } else if ha < hb {
        uf.parents[ra] = rb
    } else {
        // ha == hb
        uf.parents[ra] = rb
        uf.heights[rb] = hb + 1
    }

    uf.count--
}

func isSimilar(a, b string) bool {
    n := len(a)
    diff := 0
    for i := 0; i < n; i++ {
        if a[i] != b[i] {
            diff++
            if diff > 2 {
                return false
            }
        }
    }
    return diff == 2
}

func numSimilarGroups(A []string) int {
    setA := make(map[string]bool)
    for _, s := range A {
        setA[s] = true
    }
    A = A[0:0]
    for s, _ := range setA {
        A = append(A, s)
    }

    n := len(A)
    uf := newUF(n)

    for i := 0; i < n; i++ {
        for j := i + 1; j < n; j++ {
            if isSimilar(A[i], A[j]) {
                uf.union(i, j)
            }
        }
    }

    return uf.count
}