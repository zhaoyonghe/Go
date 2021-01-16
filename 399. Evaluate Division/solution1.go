func put(m map[string]map[string]float64, s1 string, s2 string, val float64) {
    mm, has := m[s1]
    if has {
        mm[s2] = val;
        m[s1] = mm;
    } else {
        tmp := make(map[string]float64)
        tmp[s2] = val
        m[s1] = tmp 
    }
}

func dfs(g map[string]map[string]float64, start string, target string, status map[string]int, curVal float64) (float64, bool) {
    if start == target {
        return curVal, true
    }

    status[start] = 1

    neighbors := g[start]
    for next, v := range neighbors {
        if status[next] == 0 {
            // step to node next
            res, find := dfs(g, next, target, status, curVal * v)
            if find {
                return res, find
            }
        }
    }

    status[start] = 2
    return -1.0, false
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
    g := make(map[string]map[string]float64)

    n := len(values)

    for i := 0; i < n; i++ {
        put(g, equations[i][0], equations[i][1], values[i])
        put(g, equations[i][1], equations[i][0], 1.0 / values[i])
    }

    result := []float64{}

    for _, q := range queries {
        _, has0 := g[q[0]]
        _, has1 := g[q[1]]
        if !has0 || !has1 {
            result = append(result, -1.0)
            continue
        }
        res, _ := dfs(g, q[0], q[1], make(map[string]int), 1.0)
        result = append(result, res)
    }

    return result
}