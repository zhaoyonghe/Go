func strangePrinter(s string) int {
    dp := make(map[[2]int]int)

    news := []byte{}
    
    for i := 0; i < len(s); i++ {
        if i == 0 || s[i] != s[i - 1] {
            news = append(news, s[i])
        }
    }
    //fmt.Printf("%v\n", news)

    return dfs(news, 0, len(news) - 1, dp)
}

func dfs(s []byte, i int, j int, dp map[[2]int]int) int {
    if i > j {
        return 0
    }
    if i == j {
        return 1
    }
    if val, ok := dp[[2]int{i, j}]; ok {
        return val
    }

    res := math.MaxInt32
    res = min(res, dfs(s, i, j - 1, dp) + 1)

    for k := j - 1; k >= i; k-- {
        if s[k] == s[j] {
            res = min(res, dfs(s, i, k, dp) + dfs(s, k + 1, j - 1, dp))
        }
    }

    dp[[2]int{i, j}] = res
    return res
}

func min(a, b int) int {
    if a < b {
        return a
    } else {
        return b
    }
}