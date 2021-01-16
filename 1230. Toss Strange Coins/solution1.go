func probabilityOfHeads(prob []float64, target int) float64 {
    n := len(prob)
    dp := make([][]float64, n)
    for i := 0; i < n; i++ {
        dp[i] = make([]float64, target + 1)
    }

    dp[0][0] = 1 - prob[0]
    for i := 1; i < n; i++ {
        dp[i][0] = dp[i - 1][0] * (1 - prob[i])
    }

    if target == 0 {
        return dp[n - 1][0]
    }

    dp[0][1] = prob[0]

    for i := 1; i < n; i++ {
        for j := 1; j <= target; j++ {
            dp[i][j] = prob[i] * dp[i - 1][j - 1] + (1 - prob[i]) * dp[i - 1][j]
        }
    }

    return dp[n - 1][target]
}