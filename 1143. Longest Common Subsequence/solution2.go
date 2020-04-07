func max(a int, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}

func longestCommonSubsequence(text1 string, text2 string) int {
    var l1 int = len(text1)
    var l2 int = len(text2)

    var dp []int = make([]int, l2 + 1)
    var luc int = 0

    for i := 1; i <= l1; i++ {
        luc = dp[0]
        for j := 1; j <= l2; j++ {
            if text1[i - 1] == text2[j - 1] {
                dp[j], luc = luc + 1, dp[j]
            } else {
                dp[j], luc = max(dp[j], dp[j - 1]), dp[j]
            }
        }
        // fmt.Printf("%v\n", dp)
    }

    return dp[l2]
}