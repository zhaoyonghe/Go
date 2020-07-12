func calculateMinimumHP(dungeon [][]int) int {
    m := len(dungeon)
    n := len(dungeon[0])
    dp := make([]int, n)
    for i := range dp {
        dp[i] = 1 << 30
    }
    dp[n - 1] = 1

    for i := m - 1; i >= 0; i-- {
        dp[n - 1] = max(1, dp[n - 1] - dungeon[i][n - 1])
        for j := n - 2; j >= 0; j-- {
            dp[j] = min(max(1, dp[j] - dungeon[i][j]), max(1, dp[j + 1] - dungeon[i][j]))
        } 
    }

    return dp[0]
}

func min(a, b int) int {
    if a > b {
        return b
    }
    return a
}

func max(a, b int) int {
    if a < b {
        return b
    }
    return a
}