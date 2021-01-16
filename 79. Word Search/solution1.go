func exist(board [][]byte, word string) bool {
    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[0]); j++ {
            if dfs(board, word, i, j, 0) {
                return true
            }
        }
    }
    return false
}

func dfs(board [][]byte, word string, i, j, cur int) bool {
    if cur == len(word) {
        return true
    }
    if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) || board[i][j] != word[cur] {
        return false
    }

    c := board[i][j]
    board[i][j] = '*'
    res := dfs(board, word, i + 1, j, cur + 1) || dfs(board, word, i - 1, j, cur + 1) || dfs(board, word, i, j + 1, cur + 1) || dfs(board, word, i, j - 1, cur + 1)
    board[i][j] = c

    return res
}