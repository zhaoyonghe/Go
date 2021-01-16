type Corner struct {
    x int
    h int
    iss bool
    hidx int
}

type CornerList []*Corner

func (carr CornerList) Len() int {
    return len(carr)
}

func (carr CornerList) Less(i int, j int) bool {
    if carr[i].x != carr[j].x {
        return carr[i].x < carr[j].x
    }

    if carr[i].iss && carr[j].iss {
        return carr[i].h > carr[j].h
    }

    if !carr[i].iss && !carr[j].iss {
        return carr[i].h < carr[j].h
    }

    return carr[i].iss
}

func (carr CornerList) Swap(i int, j int) {
    carr[i], carr[j] = carr[j], carr[i]
}

type maxHeap []*Corner

func (h maxHeap) Len() int {
    return len(h)
}

func (h maxHeap) Less(i, j int) bool {
    return h[i].h > h[j].h
}

func (h maxHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
    h[i].hidx = i
    h[j].hidx = j
}

func (h *maxHeap) Push(x interface{}) {
    c := x.(*Corner)
    c.hidx = len(*h)
    *h = append(*h, c)
} 

func (h *maxHeap) Pop() interface{} {
    old := *h
    n := len(old)
    res := old[n - 1]
    *h = old[:n - 1]
    res.hidx = -34
    return res
}


func getSkyline(buildings [][]int) [][]int {
    carr := make([]*Corner, 0, len(buildings) * 2)
    m := make(map[*Corner]*Corner)

    for _, b := range buildings {
        left := &Corner{b[0], b[2], true, -1}
        right := &Corner{b[1], b[2], false, -1}
        carr = append(carr, left)
        carr = append(carr, right)
        m[right] = left
    }

    sort.Sort(CornerList(carr))

    res := make([][]int, 0, 1)
    h := &maxHeap{}
    heap.Init(h) 

    for _, c := range carr {
        if c.iss {
            if (*h).Len() == 0 {
                res = append(res, []int{c.x, c.h})
            } else {
                if c.h > (*h)[0].h {
                    res = append(res, []int{c.x, c.h})
                }
            }
            heap.Push(h, c)
        } else {
            before := (*h)[0]

            heap.Remove(h, (m[c]).hidx)

            if (*h).Len() == 0 {
                res = append(res, []int{c.x, 0})
            } else {
                now := (*h)[0]
                if before.h > now.h {
                    res = append(res, []int{c.x, now.h})
                }
            }
        }
    } 

    return res
}