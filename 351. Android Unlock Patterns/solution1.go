func numberOfPatterns(m int, n int) int {
    visited := [9]bool{}
    res := [9]int{}

    dfs(0, &visited, 0, &res, n)
    dfs(1, &visited, 0, &res, n)
    for i := 0; i < 9; i++ {
        res[i] *= 4
    }

    dfs(4, &visited, 0, &res, n)

    sum := 0
    for i := m - 1; i <= n - 1; i++ {
        sum += res[i]
    }
    return sum
}

func dfs(cur int, visited *[9]bool, count int, res *[9]int, ub int) {
    if count == ub {
        return
    }
    visited[cur] = true
    (*res)[count]++
    for i := 0; i < 9; i++ {
        if isValid(cur, i, visited) {
            dfs(i, visited, count + 1, res, ub)
        }
    }
    visited[cur] = false
}

func isValid(from int, to int, visited *[9]bool) bool {
    if (*visited)[to] {
        return false
    } else {
        xdiff := abs(from % 3 - to % 3)
        ydiff := abs(from / 3 - to / 3)

        if xdiff == 2 && ydiff == 2 {
            return (*visited)[4]
        }
        if xdiff == 2 && ydiff == 0 {
            if from == 0 || from == 2 {
                return (*visited)[1]
            }
            if from == 3 || from == 5 {
                return (*visited)[4]
            }
            if from == 6 || from == 8 {
                return (*visited)[7]
            }
        }
        if xdiff == 0 && ydiff == 2 {
            if from == 0 || from == 6 {
                return (*visited)[3]
            }
            if from == 1 || from == 7 {
                return (*visited)[4]
            }    
            if from == 2 || from == 8 {
                return (*visited)[5]
            }          
        }
        return true
    }
}

func abs(a int) int {
    if a < 0 {
        return -a
    } else {
        return a
    }
}