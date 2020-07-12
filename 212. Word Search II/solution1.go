type Trie struct {
    str string
    next [26]*Trie
}

func indexOf(c byte) int {
    return int(c - 'a')
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
    cur := this
    for i := 0; i < len(word); i++ {
        c := word[i]
        if cur.next[indexOf(c)] == nil {
            cur.next[indexOf(c)] = &Trie{}
        }
        cur = cur.next[indexOf(c)]
    }
    cur.str = word
}

func findWords(board [][]byte, words []string) []string {
    root := &Trie{}
    for _, w := range words {
        root.Insert(w)
    }

    res := &[]string{}
    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[0]); j++ {
            dfs(board, i, j, root, res)
        }
    }
    sort.Strings(*res)
    return *res
}

func dfs(board [][]byte, i, j int, cur *Trie, res *[]string) {
    if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) || board[i][j] == '*' || cur.next[indexOf(board[i][j])] == nil {
        return
    }

    tmp := cur.next[indexOf(board[i][j])]
    if tmp.str != "" {
        *res = append(*res, tmp.str)
        tmp.str = ""
    }

    
    c := board[i][j]
    board[i][j] = '*'
    dfs(board, i + 1, j, tmp, res)
    dfs(board, i - 1, j, tmp, res)
    dfs(board, i, j + 1, tmp, res)
    dfs(board, i, j - 1, tmp, res)
    board[i][j] = c
}