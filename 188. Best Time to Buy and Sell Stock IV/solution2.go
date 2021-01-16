func maxProfit(k int, prices []int) int {
    n := len(prices)

    if n <= 1 || k < 1 {
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

    hold := make([][]int, 2)
    unhold := make([][]int, 2)

    hold[0] = make([]int, n)
    hold[1] = make([]int, n)
    unhold[0] = make([]int, n)
    unhold[1] = make([]int, n)

    hold[0][0] = -prices[0]
    hold[1][0] = -prices[0]
    for i := 1; i < n; i++ {
        hold[0][i] = max(hold[0][i - 1], -prices[i])
    }

    for tran := 1; tran <= k; tran++ {
        for day := 1; day < n; day++ {
            hold[1][day] = max(unhold[1][day - 1] - prices[day], hold[1][day - 1])
            unhold[1][day] = max(hold[0][day - 1] + prices[day], unhold[1][day - 1])
        }
        if unhold[0][n - 1] == unhold[1][n - 1] {
            return unhold[1][n - 1]
        }
        hold[0], hold[1] = hold[1], hold[0]
        unhold[0], unhold[1] = unhold[1], unhold[0]
        //fmt.Printf("=======================\n")
        //fmt.Printf("%v\n", hold[0])
        //fmt.Printf("%v\n", hold[1])
        //fmt.Printf("%v\n", unhold[0])
        //fmt.Printf("%v\n", unhold[1])
    }

    return unhold[0][n - 1]
}

func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}