func lastStoneWeightII(stones []int) int {
    sum := 0
    for _, val := range stones {
        sum += val
    }

    dp := make([]bool, sum / 2 + 1)
    dp[0] = true

    for _, val := range stones {
        for j := sum / 2; j >= val; j-- {
            dp[j] = dp[j] || dp[j - val]
        }
    }

    for j := sum / 2; j >= 0; j-- {
        if dp[j] {
            return sum - j - j
        }
    }

    return -1
}