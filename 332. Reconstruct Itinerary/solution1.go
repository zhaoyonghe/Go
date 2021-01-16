type StringHeap []string
 
func (h StringHeap) Len() int { 
    return len(h) 
}

func (h StringHeap) Less(i, j int) bool { 
    return h[i] < h[j]
}

func (h StringHeap) Swap(i, j int) { 
    h[i], h[j] = h[j], h[i] 
}

func (h *StringHeap) Push(x interface{}) {
    *h = append(*h, x.(string))
}

func (h *StringHeap) Pop() interface{} {
	n := len(*h)
	x := (*h)[n - 1]
	*h = (*h)[:n-1]
    return x
}

func reverse(res []string) []string {
    n := len(res)
    for i := 0; i < n / 2; i++ {
        res[i], res[n - i - 1] = res[n - i - 1], res[i]
    }
    return res
}

type Graph map[string]*StringHeap
func (g Graph) newIfAbsent(key string) {
    if _, has := g[key]; !has {
        h := &StringHeap{}
        heap.Init(h)
        g[key] = h
    }
}

func dfs(cur string, g Graph, res *[]string) {
    h := g[cur]
    for (*h).Len() > 0 {
        next := heap.Pop(h).(string)
        dfs(next, g, res)
    }
    *res = append(*res, cur)
}

func findItinerary(tickets [][]string) []string {
    g := (Graph)(make(map[string]*StringHeap))

    for _, ticket := range tickets {
        g.newIfAbsent(ticket[0])
        g.newIfAbsent(ticket[1])
        heap.Push(g[ticket[0]], ticket[1])
    }
    
    res := []string{}
    dfs("JFK", g, &res)

    return reverse(res)
}