func maxProfit(prices []int) int {
    if len(prices) == 0 || len(prices) == 1 {
        return 0
    }
    
    bottom := prices[0]
    peak := prices[0]
    sum := 0
    for _, val := range prices {
        if val < bottom {
            sum = max(sum, peak - bottom)
            bottom = val
            peak = val
        }
        if val > peak {
            peak = val
        }
    }

    sum = max(sum, peak - bottom)
    return sum
}

func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}