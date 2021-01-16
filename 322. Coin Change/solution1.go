func min(a, b int) int {
    if (a > b) {
        return b
    } else {
        return a
    }
}

func coinChange(coins []int, amount int) int {
    dp := make([]int, amount + 1)

    for i := 1; i <= amount; i++ {
        dp[i] = math.MaxInt32;
    }

    for i := 1; i <= amount; i++ {
        for _, c := range coins {
            if i - c >= 0 && dp[i - c] != math.MaxInt32 {
                dp[i] = min(dp[i], dp[i - c] + 1)
            }
        }
    }

    if dp[amount] == math.MaxInt32 {
        return -1
    }
    return dp[amount]  
}