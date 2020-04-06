type tdarr [][]int

func (p tdarr) Len() int {
    return len(p)
}

func (p tdarr) Less(i, j int) bool {
    return p[i][0] < p[j][0]
}

func (p tdarr) Swap(i, j int) {
    p[i], p[j] = p[j], p[i]
}

func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}


func merge(intervals [][]int) [][]int {
    var res [][]int

    sort.Sort(tdarr(intervals))

    var j int = -1

    for i := 0; i < len(intervals); i++ {
        if len(res) == 0 || res[j][1] < intervals[i][0] {
            res = append(res, []int{intervals[i][0], intervals[i][1]})
            j++
        } else {
            res[j][1] = max(res[j][1], intervals[i][1])
        }
    }

    return res
}

func insert(intervals [][]int, newInterval []int) [][]int {
    return merge(append(intervals, newInterval))
}