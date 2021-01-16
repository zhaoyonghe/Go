type queue []int

func newQueue() *queue {
    a := (queue)([]int{})
    return &a
}

func (q *queue) size() int {
    return len(*q)
}

func (q *queue) offer(a int) {
    *q = append(*q, a)
}

func (q *queue) poll() int {
    a := (*q)[0]
    *q = (*q)[1:]
    return a
}

func findMinHeightTrees(n int, edges [][]int) []int {
    if n == 1 {
        return []int{0}
    }
    degrees := make([]int, n)
    connect := make([]int, n)
    q := newQueue()

    for _, e := range edges {
        degrees[e[0]]++
        degrees[e[1]]++

        connect[e[0]] ^= e[1]
        connect[e[1]] ^= e[0]
    }

    for i, val := range degrees {
        if val == 1 {
            q.offer(i)
        }
    }

    for n > 2 {
        sz := q.size()
        for i := 0; i < sz; i++ {
            leaf := q.poll()
            
            to := connect[leaf]

            degrees[to]--
            if degrees[to] == 1 {
                q.offer(to)
            }
            // remove the connection to the leaf
            connect[to] ^= leaf
        }
        n -= sz
    }
    return ([]int)(*q)
}