type Item struct {
    id int
    time int
    hidx int
}

type ItemList []*Item

func (il ItemList) Len() int {
    return len(il)
}

func (il ItemList) Less(i, j int) bool {
    return il[i].time < il[j].time
}

func (il ItemList) Swap(i, j int) {
    il[i], il[j] = il[j], il[i]
    il[i].hidx = i
    il[j].hidx = j
}

func (il *ItemList) Pop() interface{} {
    old := *il
    n := len(old)
    item := old[n - 1]
    *il = old[0:n - 1]
    item.hidx = -1
    return item
}

func (il *ItemList) Push(x interface{}) {
    item := x.(*Item)
    item.hidx = len(*il)
    *il = append(*il, item)
}

func networkDelayTime(times [][]int, N int, K int) int {
    g := make([]map[int]int, N)
    for i := 0; i < N; i++ {
        g[i] = make(map[int]int)
    }

    for _, val := range times {
        g[val[0] - 1][val[1] - 1] = val[2]
    }

    record := make([]*Item, N)

    pq := &ItemList{}
    heap.Init(pq)

    // insert the items into the pq
    for i := 0; i < N; i++ {
        if i == K - 1 {
            record[i] = &Item{i, 0, -1}
        } else {
            record[i] = &Item{i, math.MaxInt32, -1}
        }
        heap.Push(pq, record[i])
    }

    res := 0
    for len(*pq) != 0 {
        // the path length from K to node item.id is determined
        item := heap.Pop(pq).(*Item)
        res = item.time
        neighbors := g[item.id]
        for neiid, w := range neighbors {
            if record[neiid].hidx >= 0 {
                // neiid is still in the heap
                neiitem := heap.Remove(pq, record[neiid].hidx).(*Item)
                if neiitem.time > item.time + w {
                    neiitem.time = item.time + w
                }
                heap.Push(pq, neiitem)
            }
        }
    }

    if res == math.MaxInt32 {
        return -1
    }

    return res


}