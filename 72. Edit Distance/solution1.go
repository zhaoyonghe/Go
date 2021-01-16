func min(a int, b int) int {
    if a < b {
        return a
    } else {
        return b
    }
}

func minDistance(word1 string, word2 string) int {
    var l1 int = len(word1)
    var l2 int = len(word2)

    if l2 == 0 {
        return l1
    }

    if l1 == 0 {
        return l2
    }

    var dp []int = make([]int, l2 + 1)
    for j := 0; j <= l2; j++ {
        dp[j] = j
    }
    var luc int = 0

    for i := 1; i <= l1; i++ {
        dp[0] = i - 1
        luc = dp[0]
        for j := 1; j <= l2; j++ {
            if word1[i - 1] == word2[j - 1] {
                dp[j], luc = min(min(dp[j - 1] + 1, dp[j] + 1), luc), dp[j]
            } else {
                dp[j], luc = min(min(dp[j - 1], dp[j]), luc) + 1, dp[j]
            }
        }
    }

    return dp[l2]
}