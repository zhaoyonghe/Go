

func maxSumAfterPartitioning(A []int, K int) int {
    dp := make([]int, len(A) + 1)
    dp[len(A)] = 0
    for i := 0; i < len(A); i++ {
        dp[i] = math.MinInt32
    }
    return dfs(A, 0, K, dp)
}

func dfs(A []int, sf int, K int, dp []int) int {
    if sf == len(A) {
        return 0
    }
    res := math.MinInt32
    ma := A[sf]

    for i := sf; i < len(A) && i < sf + K; i++ {
        ma = max(ma, A[i])
        if dp[i + 1] != math.MinInt32 {
            res = max(res, ma * (i - sf + 1) + dp[i + 1])
        } else {
            res = max(res, ma * (i - sf + 1) + dfs(A, i + 1, K, dp))
        }
    }

    dp[sf] = res
    //fmt.Printf("%v\n", dp)
    return res
}

func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}