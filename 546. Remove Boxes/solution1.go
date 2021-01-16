func removeBoxes(boxes []int) int {
    dp := make(map[[3]int]int)
    return dfs(boxes, 0, len(boxes) - 1, 0, dp)
}

func dfs(boxes []int, i int, j int, tail int, dp map[[3]int]int) int {
    //fmt.Printf("%v, %v, %v\n", i, j, tail)
    if i > j {
        return 0
    }
    if i == j {
        return (tail + 1) * (tail + 1)
    }
    if val, ok := dp[[3]int{i, j, tail}]; ok {
        return val
    }

    res := 0
    if val, ok := dp[[3]int{i, j - 1, 0}]; ok {
        res = max(res, val + (tail + 1) * (tail + 1))
    } else {
        res = max(res, dfs(boxes, i, j - 1, 0, dp) + (tail + 1) * (tail + 1))
    }

    for k := j - 1; k >= i; k-- {
        if boxes[k] == boxes[j] {
            tmp := 0
            if val, ok := dp[[3]int{i, k, tail + 1}]; ok {
                tmp += val
            } else {
                tmp += dfs(boxes, i, k, tail + 1, dp)
            }
            if val, ok := dp[[3]int{k + 1, j - 1, 0}]; ok {
                tmp += val
            } else {
                tmp += dfs(boxes, k + 1, j - 1, 0, dp)
            }
            res = max(res, tmp)
        }
    }

    dp[[3]int{i, j, tail}] = res
    return res
}

func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}