func minimumMoves(arr []int) int {
    n := len(arr)

    dp := make([][]int, n)
    for i := 0; i < n; i++ {
        dp[i] = make([]int, n)
    }

    for i := 0; i < n; i++ {
        dp[i][i] = 1
    }

    for i := 1; i < n; i++ {
        if arr[i] == arr[i - 1] {
            dp[i - 1][i] = 1
        } else {
            dp[i - 1][i] = 2
        }
    }

    for l := 2; l < n; l++ {
        for i := 0; i < n - l; i++ {
            j := i + l
            if arr[i] == arr[j] {
                dp[i][j] = dp[i + 1][j - 1]
            } else {
                dp[i][j] = math.MaxInt32
            }
            for k := i; k < j; k++ {
                dp[i][j] = min(dp[i][j], dp[i][k] + dp[k + 1][j])
            }
        }
    }

    return dp[0][n - 1]
}

func min(a, b int) int {
    if a > b {
        return b
    } else {
        return a
    }
}