func get99() *[9][9]bool {
    var m [9][9]bool
    return &m
}

func solveSudoku(board [][]byte)  {
    rows, cols, boxes := get99(), get99(), get99()
    for i := 0; i < 9; i++ {
        for j := 0; j < 9; j++ {
            // fmt.Printf("%v, %v\n", i, j)
            if board[i][j] != '.' {
                d := (int)(board[i][j] - '1')
                rows[i][d] = true
                cols[j][d] = true
                boxes[(i / 3) * 3 + j / 3][d] = true
            }
        }
    }

    solve(&board, rows, cols, boxes, 0)
}

func solve(board *[][]byte, rows *[9][9]bool, cols *[9][9]bool, boxes *[9][9]bool, k int) bool {
    if k == 81 {
        return true
    }

    var i int = k / 9
    var j int = k % 9

    if (*board)[i][j] == '.' {
        for d := 0; d < 9; d++ {
            if !rows[i][d] && !cols[j][d] && !boxes[(i / 3) * 3 + j / 3][d] {
                (*board)[i][j] = (byte)((int)('1') + d)
                rows[i][d] = true
                cols[j][d] = true
                boxes[(i / 3) * 3 + j / 3][d] = true
                if solve(board, rows, cols, boxes, k + 1) {
                    return true
                } else {
                    (*board)[i][j] = '.'
                    rows[i][d] = false
                    cols[j][d] = false
                    boxes[(i / 3) * 3 + j / 3][d] = false
                }
            }
        }
        return false
    } else {
        return solve(board, rows, cols, boxes, k + 1)
    }
}