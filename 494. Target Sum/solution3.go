func findTargetSumWays(nums []int, S int) int {
    sumall := 0
    for _, val := range nums {
        sumall += val
    }

    if sumall < abs(S) {
        return 0
    }

    if (S + sumall) % 2 != 0 {
        return 0
    }

    target := (S + sumall) / 2

    dp := make([]int, target + 1)
    dp[0] = 1
    for _, num := range nums {
        for j := target; j >= num; j-- {
            dp[j] += dp[j - num]
        }
    }

    return dp[target]
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}