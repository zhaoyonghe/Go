func stoneGameII(piles []int) int {
    sum := 0
    for _, val := range piles {
        sum += val
    }

    return (sum + dfs(piles, 0, 1, make(map[[2]int]int))) / 2
}

func dfs(piles []int, i int, m int, memo map[[2]int]int) int {
    if i == len(piles) {
        return 0
    }
    // if can take all
    if len(piles) - i < 2 * m {
        sum := 0
        for _, val := range piles[i:] {
            sum += val
        }
        return sum
    }
    // cache hit
    if val, has := memo[[2]int{i, m}]; has {
        return val
    }
    
    best := math.MinInt32
    sum := 0
    for j := i; j - i + 1 <= 2 * m && j < len(piles); j++ {
        sum += piles[j]
        newm := max(m, j - i + 1)
        newm = min(newm, (len(piles) - j) / 2)
        best = max(best, sum - dfs(piles, j + 1, newm, memo))
    }

    memo[[2]int{i, m}] = best
    return best
}

func min(a, b int) int {
    if a < b {
        return a
    } else {
        return b
    }
}

func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}