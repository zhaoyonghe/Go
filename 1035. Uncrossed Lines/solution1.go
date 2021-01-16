func maxUncrossedLines(A []int, B []int) int {
    var table [][]int = make([][]int, len(A))
    for i := 0; i < len(A); i++ {
        table[i] = make([]int, len(B))
    }

    for i := 0; i < len(A); i++ {
        for j := 0; j < len(B); j++ {
            table[i][j] = -1
        }
    }

    return dp(A, B, 0, 0, table)
}

func max(a int, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}

func dp(A []int, B []int, i int, j int, table [][]int) int {
    if i == len(A) || j == len(B) {
        return 0
    }

    if table[i][j] != -1 {
        return table[i][j]
    }

    if A[i] == B[j] {
        table[i][j] = 1 + dp(A, B, i + 1, j + 1, table)
    } else {
        table[i][j] = max(dp(A, B, i, j + 1, table), dp(A, B, i + 1, j, table))
    }

    return table[i][j]
}