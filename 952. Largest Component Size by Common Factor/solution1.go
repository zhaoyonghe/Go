type UF struct {
    parents []int
    heights []int
}

func newUF(n int) *UF {
    m1 := make([]int, n + 1)
    m2 := make([]int, n + 1)

    for i := 0; i <= n; i++ {
        m1[i] = i
        m2[i] = 0
    }

    return &UF {
        m1,
        m2,
    }
}

func (uf *UF) root(i int) int {
    for {
        ii := uf.parents[i]
        if ii == i {
            return i
        }
        i = ii
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

    if ha < hb {
        uf.parents[ra] = rb
    } else if ha > hb {
        uf.parents[rb] = ra
    } else {
        uf.parents[rb] = ra
        uf.heights[ra]++
    }
}

func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}

func maxElement(a []int) int {
    res := -1
    for _, val := range a {
        res = max(res, val)
    }
    return res
}

func largestComponentSize(A []int) int {
    n := maxElement(A)

    uf := newUF(n)

    for _, val := range A {
        for i := 2; i <= int(math.Sqrt(float64(val))); i++ {
            if val % i == 0 {
                uf.union(val, i)
                uf.union(val, val / i)
            }
        }
    }

    m := make(map[int]int)
    res := 0

    for _, val := range A {
        rt := uf.root(val)
        c, has := m[rt]
        if !has {
            m[rt] = 1
            res = max(res, 1)
        } else {
            m[rt] = 1 + c
            res = max(res, 1 + c)
        }
    }

    return res
}