func minScoreTriangulation(A []int) int {
    /*
    dp := make([][]int, len(A))
    for i := 0; i < len(A); i++ {
        dp[i] = make([]int, len(A))
    }*/
    dp := [50][50]int{}

    return dfs(A, 0, len(A) - 1, &dp)
}

func dfs(A []int, i int, j int, dp *[50][50]int) int {
    if j - i == 1 {
        return 0
    }
    if j - i == 2 {
        return A[i] * A[i + 1] * A[i + 2]
    }
    if (*dp)[i][j] != 0 {
        return (*dp)[i][j]
    }

    res := math.MaxInt32

    mul := A[i] * A[j]

    for k := i + 1; k < j; k++ {
        res = min(res, mul * A[k] + dfs(A, i, k, dp) + dfs(A, k, j, dp))
    }

    (*dp)[i][j] = res
    return res
}

func min(a, b int) int {
    if a < b {
        return a
    } else {
        return b
    }
}