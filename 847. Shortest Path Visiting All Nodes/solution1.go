type State struct {
    cur int
    vis int
}

type queue []State

func (q *queue) Push(s State) {
    *q = append(*q, s)
}

func (q *queue) Poll() State {
    res := (*q)[0]
    *q = (*q)[1:]
    return res
}

func shortestPathLength(graph [][]int) int {
    q := &queue{}

    pl := 0
    n := len(graph)

    target := (1 << n) - 1
    m := make(map[State]bool)

    for i := 0; i < n; i++ {
        s := State{i, 1 << i}
        q.Push(s)
        m[s] = true
    }

    for {
        sz := len(*q)
        for i := 0; i < sz; i++ {
            s := q.Poll()
            if target == s.vis {
                return pl
            } else {
                for _, nei := range graph[s.cur] {
                    s := State{nei, s.vis | (1 << nei)}
                    if  _, has := m[s]; !has {
                        q.Push(s)
                        m[s] = true
                    } 
                }
            }
        }
        pl++
    }

    return -1
}