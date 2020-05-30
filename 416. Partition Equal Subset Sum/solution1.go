func canPartition(nums []int) bool {
    sumall := 0
    for _, val := range nums {
        sumall += val
    }

    if sumall % 2 != 0 {
        return false
    }

    target := sumall / 2
    dp := make([]bool, target + 1)
    dp[0] = true
    fmt.Printf("%v\n", nums)

    for _, num := range nums {
        for j := target; j >= num; j-- {
            dp[j] = dp[j] || dp[j - num]
        }
        //fmt.Printf("%v\n", dp)
    }

    return dp[target]
}