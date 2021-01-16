var sumall int

func findTargetSumWays(nums []int, S int) int {
    for _, val := range nums {
        sumall += val
    }

    n := len(nums)

    dp := make([][]int, n + 1)
    for i := 0 ; i < n; i++ {
        dp[i] = make([]int, 2 * sumall + 1)
        for j := 0; j <= 2 * sumall; j++ {
            dp[i][j] = -1
        }
    }

    dp[n] = make([]int, 2 * sumall + 1)
    for j := 0; j <= 2 * sumall; j++ {
        dp[n][j] = 0
    }
    dp[n][sumall] = 1
    
    return dfs(nums, 0, S + sumall, dp)
}

func dfs(nums []int, i int, target int, dp [][]int) int {
    if target < 0 || target > 2 * sumall {
        return 0
    }
    
    a := target + nums[i]
    b := target - nums[i]
    res := 0

    if a >= 0 && a <= 2 * sumall {
        if dp[i + 1][a] == -1 {
            dp[i + 1][a] = dfs(nums, i + 1, a, dp)
        }
        res += dp[i + 1][a]
    }
    if b >= 0 && b <= 2 * sumall {
        if dp[i + 1][b] == -1 {
            dp[i + 1][b] = dfs(nums, i + 1, b, dp)
        }
        res += dp[i + 1][b]
    }
    
    dp[i][target] = res
    return res
}