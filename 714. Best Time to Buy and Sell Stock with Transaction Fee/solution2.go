func maxProfit(prices []int, fee int) int {
    n := len(prices) + 1
    hold := math.MinInt32
    unhold := 0

    for i := 1; i < n; i++ {
        hold, unhold = max(hold, unhold - prices[i - 1]), max(unhold, hold + prices[i - 1] - fee)
    }

    return max(hold, unhold)
}

func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}