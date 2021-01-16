func max(a int, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}

func longestValidParentheses(s string) int {
    if len(s) <= 1 {
        return 0
    }
    var dp []int = make([]int, len(s) + 1)
    var res int = 0

    for i := 1; i < len(s); i++ {
        var j int = i + 1
        if s[i] == ')' && s[i - 1] == '(' {
            dp[j] = dp[j - 2] + 2
        }
        if s[i] == ')' && s[i - 1] == ')' {
            if temp := i - 1 - dp[j - 1]; temp >= 0 && s[temp] == '(' {
                // dp[j] = dp[j - 1 - dp[j - 1] - 1] + dp[j - 1] + 2
                dp[j] = dp[temp] + dp[j - 1] + 2
            }
        }
        res = max(res, dp[j])
    }
    // fmt.Printf("%v\n", dp)

    return res
}