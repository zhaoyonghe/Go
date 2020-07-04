func maxProfit(prices []int) int {
    bottom := prices[0]
    peak := prices[0]
    sum := 0
    for _, val := range prices {
        if val < peak {
            sum += (peak - bottom)
            bottom = val
            peak = val
        } else {
            peak = val
        }
    }

    sum += (peak - bottom)
    return sum
}