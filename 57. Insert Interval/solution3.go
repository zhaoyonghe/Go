func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}

func min(a, b int) int {
    if a < b {
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
    var res [][]int

    var n int = len(intervals)
    var left int = newInterval[0]
    var right int = newInterval[1]
    var flag bool = true

    for i := 0; i < n; i++ {
        if left <= intervals[i][1] && intervals[i][0] <= right {
            // overlap
            left = min(left, intervals[i][0])
            right  = max(right, intervals[i][1])
        } else {
            if right < intervals[i][0] {
                res = append(append(res, []int{left, right}), intervals[i:]...)
                flag = false
                break
            } else {
                // left > intervals[i][1]
                res = append(res, intervals[i])
            }
        }
    }

    if flag {
        return append(res, []int{left, right})
    } else {
        return res
    }
}