func maxProfit(prices []int, fee int) int {
    n := len(prices) + 1
    hold := make([]int, n)
    unhold := make([]int, n)

    hold[0] = math.MinInt32

    for i := 1; i < n; i++ {
        hold[i] = max(hold[i - 1], unhold[i - 1] - prices[i - 1])
        unhold[i] = max(unhold[i - 1], hold[i - 1] + prices[i - 1] - fee)
    }

    return max(hold[n - 1], unhold[n - 1])
}

func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}