func probabilityOfHeads(prob []float64, target int) float64 {
    if target == 0 {
        res := 1.0
        for _, p := range prob {
            res *= (1 - p)
        }
        return res
    }

    dp := make([]float64, target + 1)

    dp[0] = 1 - prob[0]
    dp[1] = prob[0]

    for _, p := range prob[1:] {
        for j := target; j >= 1; j-- {
            dp[j] = p * dp[j - 1] + (1 - p) * dp[j]
        }
        dp[0] *= (1 - p)
    }

    return dp[target]
}