func isMatch(s string, p string) bool {
    var m int = len(s) + 1
    var n int = len(p) + 1

    var dp [][]bool = make([][]bool, m)
    for i, _ := range dp {
        dp[i] = make([]bool, n)
    }

    dp[0][0] = true
    for j, c := range p {
        if c != '*' {
            break
        }
        dp[0][j + 1] = true
    }

    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            if p[j - 1] == '*' {
                dp[i][j] = dp[i - 1][j] || dp[i][j - 1]
            } else if p[j - 1] == '?' {
                dp[i][j] = dp[i - 1][j - 1]
            } else {
                dp[i][j] = dp[i - 1][j - 1] && (s[i - 1] == p[j - 1])
            }
        }
    }

    return dp[m - 1][n - 1]
}