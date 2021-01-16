type UF struct {
    roots []int
    heights []int
}

func getUF(n int) *UF {
    a := make([]int, n)
    b := make([]int, n)
    for i := 0; i < n; i++ {
        a[i] = i
    }
    return &UF{a, b}
}

func (uf *UF) find(i int) int {
    for uf.roots[i] != i {
        i = uf.roots[i]
    }
    return i
}

func (uf *UF) union(a int, b int) {
    ra := uf.find(a)
    rb := uf.find(b)

    if ra == rb {
        return 
    }

    ha := uf.heights[ra]
    hb := uf.heights[rb]

    if ha > hb {
        uf.roots[rb] = ra
    } else if ha < hb {
        uf.roots[ra] = rb
    } else {
        uf.roots[ra] = rb
        uf.heights[rb]++
    }
}

func accountsMerge(accounts [][]string) [][]string {
    emailToIdxs := make(map[string][]int)
    for i, account := range accounts {
        for _, email := range account[1:] {
            if idxs, has := emailToIdxs[email]; !has {
                emailToIdxs[email] = []int{i}
            } else {
                emailToIdxs[email] = append(idxs, i)
            }
        }
    }

    uf := getUF(len(accounts))
    for _, idxs := range emailToIdxs {
        for i := 1; i < len(idxs); i++ {
            uf.union(idxs[i], idxs[i - 1])
        } 
    }

    resm := make(map[int][]string)
    for email, idxs := range emailToIdxs {
        ruser := uf.find(idxs[0])
        if v, has := resm[ruser]; has {
            resm[ruser] = append(v, email)
        } else {
            resm[ruser] = []string{email}
        }
    }

    res := [][]string{}
    for user, emails := range resm {
        sort.Strings(emails)
        res = append(res, append([]string{accounts[user][0]}, emails...))
    }

    return res
}