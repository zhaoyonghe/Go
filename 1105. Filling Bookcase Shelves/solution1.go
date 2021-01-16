func minHeightShelves(books [][]int, shelf_width int) int {
    n := len(books)
    
    dp := make([]int, n + 1)
    // dp[n] = 0
    dp[n - 1] = books[n - 1][1]
    for i := 0; i < n - 1; i++ {
        dp[i] = math.MaxInt32
    }
    
    for i := n - 2; i >= 0; i-- {
        curHeight, curWidth := 0, 0
        for j := i; j < n; j++ {
            curWidth += books[j][0]
            if curWidth > shelf_width {
                break
            }
            curHeight = max(curHeight, books[j][1])
            dp[i] = min(dp[i], curHeight + dp[j + 1])
        }
    }
    return dp[0]
}

func max(a, b int) int {
    return int(math.Max(float64(a), float64(b)))
}

func min(a, b int) int {
    return int(math.Min(float64(a), float64(b)))
}