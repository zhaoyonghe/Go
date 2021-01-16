func removeBoxes(boxes []int) int {
    dp := make(map[[3]int]int)
    newBoxes := [][2]int{}
    i := 0
    for i < len(boxes) {
        tmp := [2]int{boxes[i], 0}
        for ; i < len(boxes) && tmp[0] == boxes[i]; i++ {
            tmp[1]++
        }
        newBoxes = append(newBoxes, tmp)
    }
    //fmt.Printf("%v\n", newBoxes)

    return dfs(newBoxes, 0, len(newBoxes) - 1, 0, dp)
}

func dfs(boxes [][2]int, i int, j int, tail int, dp map[[3]int]int) int {
    //fmt.Printf("%v, %v, %v\n", i, j, tail)
    if i > j {
        return 0
    }
    if i == j {
        return (tail + boxes[i][1]) * (tail + boxes[i][1])
    }
    if val, ok := dp[[3]int{i, j, tail}]; ok {
        return val
    }

    mt := tail + boxes[j][1]

    res := 0
    res = max(res, dfs(boxes, i, j - 1, 0, dp) + mt * mt)

    for k := j - 1; k >= i; k-- {
        if boxes[k][0] == boxes[j][0] {
            res = max(res, dfs(boxes, i, k, mt, dp) + dfs(boxes, k + 1, j - 1, 0, dp))
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