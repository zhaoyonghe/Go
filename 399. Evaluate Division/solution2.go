type tuple struct {
    s string
    v float64
}

type UF struct {
    parents map[string]tuple
    heights map[string]int
}

func newUF() *UF {
    return &UF{
        make(map[string]tuple),
        make(map[string]int),
    }
}

func (uf *UF) putIfAbsent(s string) {
    if _, has := uf.parents[s]; !has {
        uf.parents[s] = tuple{s, 1.0}
        uf.heights[s] = 1
    }
}

func (uf *UF) root(s string) (string, float64) {
    pm := 1.0
    for {
        t := uf.parents[s]
        pm *= t.v
        if s == t.s {
            return s, pm
        }
        s = t.s
    }
}

func (uf *UF) union(a string, b string, v float64) {
    ra, pma := uf.root(a)
    rb, pmb := uf.root(b)

    if ra == rb {
        return
    }

    hra := uf.heights[ra]
    hrb := uf.heights[rb]

    if hra > hrb {
        uf.parents[rb] = tuple{ra, (v * pma) / pmb}
    } else if hra < hrb {
        uf.parents[ra] = tuple{rb, pmb / (v * pma)}
    } else {
        // hra == hrb
        uf.parents[rb] = tuple{ra, (v * pma) / pmb}
        uf.heights[ra]++
    }
} 

func (uf *UF) has(s string) bool {
    _, ok := uf.parents[s]
    return ok
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
    uf := newUF()

    n := len(equations)

    for i := 0; i < n; i++ {
        a := equations[i][0]
        b := equations[i][1]
        uf.putIfAbsent(a)
        uf.putIfAbsent(b)
        uf.union(a, b, values[i])
    }

    res := []float64{}
    for _, q := range queries {
        if uf.has(q[0]) && uf.has(q[1]) {
            ra, pma := uf.root(q[0])
            rb, pmb := uf.root(q[1])
            if ra == rb {
                res = append(res, pmb / pma)
            } else {
                res = append(res, -1)
            }
        } else {
            res = append(res, -1)
        }
    }

    return res
}