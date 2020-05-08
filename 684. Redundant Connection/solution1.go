type stack []int

func newStack() *stack {
    a := (stack)([]int{})
    return &a
}

func (s *stack) push(a int) {
    *s = append(*s, a)
}

func (s *stack) pop() int {
    a := (*s)[len(*s) - 1]
    *s = (*s)[:len(*s) - 1]
    return a
}

func findRedundantConnection(edges [][]int) []int {
    n := len(edges)
    g := make([][]int, n + 1)
    for i := 0; i <= n; i++ {
        g[i] = []int{}
    }

    for i, e := range edges {
        g[e[0]] = append(g[e[0]], i)
        g[e[1]] = append(g[e[1]], i)
    }
    //for i := 0; i <= n; i++ {
    //    fmt.Printf("%v\n", g[i])
    //}    

    s := newStack()
    status := make([]int, n + 1)
    dfs(1, g, status, s, edges, -1)

    node := s.pop()

    max := s.pop()
    temp := ([]int)(*s)
    //fmt.Printf("%v\n", temp)
    for i := len(temp) - 1; i >= 0; i-- {
        val := temp[i]
        if val > max {
            max = val
        }
        if edges[val][0] == node || edges[val][1] == node {
            break
        }
    }
    return edges[max]
}

func dfs(cur int, g [][]int, status []int, s *stack, edges [][]int, prev int) bool {
    status[cur] = 1

    for _, val := range g[cur] {
        nei := getNei(edges[val], cur)
        //fmt.Printf("%v %v %v\n", prev, cur, nei)
        if status[nei] == 2 || nei == prev {
            continue
        }
        if status[nei] == 1 {
            s.push(val)
            s.push(nei)
            return true
        }
        // if status[nei] == 0
        s.push(val)
        if dfs(nei, g, status, s, edges, cur) {
            return true
        }
    }

    status[cur] = 2
    s.pop()
    return false
}

func getNei(edge []int, ex int) int {
    if edge[0] == ex {
        return edge[1]
    } else {
        return edge[0]
    }
}