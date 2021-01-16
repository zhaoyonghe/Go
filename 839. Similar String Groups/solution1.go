type UF struct {
    parents map[string]string
    heights map[string]int
}

func newUF() *UF {
    return &UF{
        make(map[string]string),
        make(map[string]int),
    }
}

func (uf *UF) putIfAbsent(a string) {
    _, has := uf.parents[a]
    if !has {
        uf.parents[a] = a
        uf.heights[a] = 0
    }
}

func (uf *UF) root(a string) string {
    for {
        p := uf.parents[a]
        if p == a {
            return p
        }
        a = p
    }
}

func (uf *UF) union(a, b string) {
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
}

func assert(b bool) {
    if !b {
        panic("dsaf")
    }
}

func isSimilar(a, b string) bool {
    assert(len(a) == len(b))

    n := len(a)
    diff := []int{}
    for i := 0; i < n; i++ {
        if a[i] != b[i] {
            diff = append(diff, i)
            if len(diff) > 2 {
                return false
            }
        }
    }
    return len(diff) == 2 && a[diff[0]] == b[diff[1]] && a[diff[1]] == b[diff[0]]
}

func numSimilarGroups(A []string) int {
    uf := newUF()

    for _, s := range A {
        uf.putIfAbsent(s)
    }

    n := len(A)
    for i := 0; i < n; i++ {
        for j := i + 1; j < n; j++ {
            if isSimilar(A[i], A[j]) {
                uf.union(A[i], A[j])
            }
        }
    }

    m := make(map[string]bool)
    for k, _ := range uf.parents {
        m[uf.root(k)] = true
    }

    return len(m)
}