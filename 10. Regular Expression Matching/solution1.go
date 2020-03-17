func isMatch(s string, p string) bool {
    var m int = len(s)
    var n int = len(p)
    var dp = make([][]bool, m+1)
    for i := 0; i <= m; i++ {
        dp[i] = make([]bool, n+1)
    }

    dp[0][0] = true
    
    for j := 1; j < n; j += 2 {
        if p[j] == '*' {
            dp[0][j + 1] = true
        } else {
            break
        }
    }
    
    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            if p[j - 1] == '*' {
                dp[i][j] = (dp[i][j - 2] || (dp[i - 1][j] && (p[j - 2] == '.' || (p[j - 2] == s[i - 1]))))
            } else if p[j - 1] == '.' {
                dp[i][j] = dp[i-1][j-1]
            } else {
                dp[i][j] = (p[j - 1] == s[i - 1]) && dp[i-1][j-1]
            }
        }
    }
    
    return dp[m][n]
}