func videoStitching(clips [][]int, T int) int {
    dp := make([][]int, T + 1)
    for i := 0; i <= T; i++ {
        dp[i] = make([]int, T + 1)
        for j := 0; j <= T; j++ {
            dp[i][j] = math.MaxInt32
            if i == j {
                dp[i][j] = 0
            }
        }
    }

    for _, c := range clips {
        s := c[0]
        e := c[1]
        for le := 1; le <= T; le++ {
            for j := le; j <= T; j++ {
                i := j - le
                if e <= i || j <= s {
                    continue
                } else if s <= i && j <= e {
                    dp[i][j] = 1
                } else if i <= s && e <= j {
                    dp[i][j] = min(dp[i][j], dp[i][s] + 1 + dp[e][j])
                } else if s <= j {
                    dp[i][j] = min(dp[i][j], dp[i][s] + 1)
                } else if i <= e{
                    dp[i][j] = min(dp[i][j], 1 + dp[e][j])
                } else {
                    panic("not happen")
                }
            } 
        }
    }

    if dp[0][T] == math.MaxInt32 {
        return -1
    } else {
        return dp[0][T]
    }
}

func min(a, b int) int {
    if a > b {
        return b
    } else {
        return a
    }
}