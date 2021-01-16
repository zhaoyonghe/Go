func get99() *[9][9]bool {
    var m [9][9]bool
    return &m
}

func ts(a *bool) bool {
    if *a {
        return true
    } else {
        *a = true
        return false
    }
}

func isValidSudoku(board [][]byte) bool {
    //var cols *[9][9]bool = get99()
    rows, cols, boxes := get99(), get99(), get99()

    for i := 0; i < 9; i++ {
        for j := 0; j < 9; j++ {
            // fmt.Printf("%v, %v\n", i, j)
            if board[i][j] != '.' {
                d := (int)(board[i][j] - '1')
                if ts(&rows[i][d]) {
                    // fmt.Printf("a\n")
                    return false
                }
                if ts(&cols[j][d]) {
                    // fmt.Printf("b\n")
                    return false
                }
                if ts(&boxes[(i / 3) * 3 + j / 3][d]) {
                    // fmt.Printf("c\n")
                    return false
                }
            }
        }
    } 
    //fmt.Printf("%v", cols[0][0])
    return true
}