func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}


func merge(intervals [][]int) [][]int {
    var res [][]int

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
    intervals = append(intervals, newInterval)
    var i int = len(intervals) - 2
    for ; i >= 0; i-- {
        if intervals[i][0] > newInterval[0] {
            intervals[i + 1] = intervals[i]
        } else {
            break
        }
    }
    intervals[i + 1] = newInterval
    return merge(intervals)
}