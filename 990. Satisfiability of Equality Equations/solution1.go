type UF struct {
    parents [26]int
    heights [26]int
}

func newUF() *UF {
    a := [26]int{}
    for i := 0; i < 26; i++ {
        a[i] = i
    }
    return &UF {
        a,
        a,
    }
}

func (uf *UF) root(i int) int {
    for {
        if uf.parents[i] == i {
            return i
        }
        i = uf.parents[i]
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
        uf.heights[rb]++
    }
}

func equationsPossible(equations []string) bool {
    uf := newUF()

    neq := [][]int{}

    for _, equation := range equations {
        le := int(equation[0] - 'a')
        ri := int(equation[3] - 'a')
        if equation[1] == '!' {
            neq = append(neq, []int{le, ri})
        } else {
            uf.union(le, ri)
        }
    }

    for _, val := range neq {
        if uf.root(val[0]) == uf.root(val[1]) {
            return false
        }
    }

    return true
}