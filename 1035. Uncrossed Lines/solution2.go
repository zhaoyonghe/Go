func maxUncrossedLines(A []int, B []int) int {
    var l1 int = len(A)
    var l2 int = len(B)

    var dp []int = make([]int, l2 + 1)
    var luc int = 0

    for i := 1; i <= l1; i++ {
        luc = dp[0]
        for j := 1; j <= l2; j++ {
            if A[i - 1] == B[j - 1] {
                dp[j], luc = luc + 1, dp[j]
            } else {
                dp[j], luc = max(dp[j], dp[j - 1]), dp[j]
            }
        }
        // fmt.Printf("%v\n", dp)
    }

    return dp[l2]
}

func max(a int, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}
