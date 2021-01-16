func maxProfit(k int, prices []int) int {
    if k <= 0 {
        return 0
    }
    if len(prices) < 2 {
        return 0
    }
    if len(prices) == 2 {
        if prices[0] < prices[1] {
            return prices[1] - prices[0]
        } else {
            return 0
        }
    }
    if k >= len(prices) / 2 {
        sum := 0
        for i := 1; i < len(prices); i++ {
            if prices[i] > prices[i - 1] {
                sum += (prices[i] - prices[i - 1])
            }
        }
        return sum
    }
    //memo := make(map[[3]int]int)
    memo := make([][][]int, len(prices))
    for i := 0; i < len(prices); i++ {
        memo[i] = make([][]int, k + 1)
        for j := 1; j <= k; j++ {
            memo[i][j] = []int{math.MinInt32, math.MinInt32}
        }
    }
    return dfs(prices, 0, 0, k, memo)
}

func dfs(prices []int, i int, isHolding int, k int, memo [][][]int) int {
    // no day
    if i == len(prices) {
        return 0
    }
    // no trans
    if k == 0 {
        return 0
    }
    // cache hit
    if memo[i][k][isHolding] != math.MinInt32 {
        return memo[i][k][isHolding]
    }

    res := -1
    if isHolding == 1 {
        // sell
        res = max(res, prices[i] + dfs(prices, i + 1, 0, k - 1, memo))
        // skip
        res = max(res, dfs(prices, i + 1, 1, k, memo))
    } else {
        // buy
        res = max(res, -prices[i] + dfs(prices, i + 1, 1, k, memo))
        // skip
        res = max(res, dfs(prices, i + 1, 0, k, memo))
    }

    memo[i][k][isHolding] = res

    return res
}

func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}