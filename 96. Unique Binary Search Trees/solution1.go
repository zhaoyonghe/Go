func numTrees(n int) int {
    if n < 3 {
        return n
    }
    dp := make([]int, n + 1)
    dp[0], dp[1], dp[2] = 1, 1, 2

    for k := 3; k <= n; k++ {
        sum := 0
        tmp := k - 1
        for i := 0; i <= tmp/2; i++ {
            sum += dp[i] * dp[tmp-i]
        }
        sum *= 2
        if k % 2 == 1 {
            sum -= dp[tmp/2] * dp[tmp/2]
        }
        dp[k] = sum
    }
    return dp[n]
}