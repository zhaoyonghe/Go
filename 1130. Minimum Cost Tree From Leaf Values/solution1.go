func mctFromLeafValues(arr []int) int {
    n := len(arr)
    dp := make([][][3]int, n)
    for i := 0; i < n; i++ {
        dp[i] = make([][3]int, n)
        //for j := 0; j < n; j++ {
        //    dp[i][j] = make([]int, 3)
        //}
    }

    for i := 0; i < n; i++ {
        dp[i][i][0], dp[i][i][1], dp[i][i][2] = arr[i], arr[i], 0
    }
    for i := 1; i < n; i++ {
        dp[i - 1][i][0], dp[i - 1][i][1] = arr[i - 1] * arr[i], max(arr[i - 1], arr[i])
        dp[i - 1][i][2] = dp[i - 1][i][0]
    }

    for sp := 2; sp < n; sp++ {
        for j := sp; j < n; j++ {
            i := j - sp
            min := math.MaxInt32
            for k := i; k < j; k++ {
                tmpsum := dp[i][k][1] * dp[k + 1][j][1] + dp[i][k][2] + dp[k + 1][j][2]
                if tmpsum < min {
                    min = tmpsum
                    dp[i][j][0], dp[i][j][1], dp[i][j][2] = dp[i][k][1] * dp[k + 1][j][1], max(dp[i][k][1], dp[k + 1][j][1]), tmpsum
                }
            }
        }
    }

    return dp[0][n - 1][2]
}

func min(a, b int) int {
    if a < b {
        return a
    } else {
        return b
    }
}

func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}