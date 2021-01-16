func maximalSquare(matrix [][]byte) int {
    m := len(matrix)
    if m == 0 {
        return 0
    }
    n := len(matrix[0])
    if n == 0 {
        return 0
    }
    dp := make([][]int, m)
    for i := 0; i < m; i++ {
        dp[i] = make([]int, n)
    }

    for i := 0; i < m; i++ {
        if matrix[i][0] == '1' {
            dp[i][0] = 1
        }
    }

    for j := 0; j < n; j++ {
        if matrix[0][j] == '1' {
            dp[0][j] = 1
        }
    }

    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            if matrix[i][j] == '1' {
                dp[i][j] = 1 + min(dp[i - 1][j], dp[i][j - 1], dp[i - 1][j - 1])
            }
        }
    }

    ma := 0
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            ma = max(dp[i][j], ma)
        }
    }

    return ma * ma
}

func max(a, b int) int {
    if a < b {
        return b
    } else {
        return a
    }
}

func min(a, b, c int) int {
    if a < b {
        if a < c {
            return a
        } else {
            return c
        }
    } else {
         if b < c {
            return b
        } else {
            return c
        }       
    }
}