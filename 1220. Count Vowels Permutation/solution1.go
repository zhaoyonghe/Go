func countVowelPermutation(n int) int {
    const d int = 1000000007
    dp := [5]int{1,1,1,1,1}
    for i := 1; i < n; i++ {
        dp[0], dp[1], dp[2], dp[3], dp[4] =
        (dp[1] + dp[2] + dp[4]) % d,
        (dp[0] + dp[2]) % d,
        (dp[1] + dp[3]) % d,
        (dp[2]) % d,
        (dp[2] + dp[3]) %d
    }

    return (dp[0] + dp[1] + dp[2] + dp[3] + dp[4]) % d
}