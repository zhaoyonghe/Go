func minFallingPathSum(A [][]int) int {
    n := len(A)
    if n == 1 {
        return A[0][0]
    }
    
    up := A[0]
    down := make([]int, n)

    for i := 1; i < n; i++ {
        down[0] = A[i][0] + min(up[0], up[1])
        down[n - 1] = A[i][n - 1] + min(up[n - 2], up[n - 1])
        for j := 1; j < n - 1; j++ {
            down[j] = A[i][j] + mines([]int{up[j - 1], up[j], up[j + 1]})
        }
        up, down = down, up
    }

    return mines(up)
}

func mines(s []int) int {
    res := math.MaxInt32
    for _, val := range s {
        res = min(res, val)
    }
    return res
}

func max(a, b int) int {
    return int(math.Max(float64(a), float64(b)))
}

func min(a, b int) int {
    return int(math.Min(float64(a), float64(b)))
}