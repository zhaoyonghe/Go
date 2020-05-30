var prime []int = []int{2,3,5,7,11,13,17,19,23,29,31,37,41,43,47,53,59,61,67,71,73,79,83,89,97,101,103,107,109,113,127,131,137,139,149,151,157,163,167,173,179,181,191,193,197,199,211,223,227,229,233,239,241,251,257,263,269,271,277,281,283,293,307,311,313,317}


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
        realVal := val
        for i := 0; i < len(prime) && prime[i] <= val; i++ {
            if val % prime[i] == 0 {
                uf.union(realVal, prime[i])
                for val % prime[i] == 0 {
                    val /= prime[i]
                }
                //uf.union(val, val / prime[i])
            }
        }
        // realVal contains > 317 prime factor
        if val != 1 {
            //fmt.Printf("%v, %v\n", realVal, val)
            uf.union(realVal, val)
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