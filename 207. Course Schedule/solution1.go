func canFinish(numCourses int, prerequisites [][]int) bool {
    var g [][]int = make([][]int, numCourses)
    var status []int = make([]int, numCourses)

    for _, pre := range prerequisites {
        g[pre[1]] = append(g[pre[1]], pre[0])
        status[pre[0]] = 0
        status[pre[1]] = 0
    }

    for k, v := range status {
        if v == 0 && !dfs(g, status, k) {
            return false
        }
    }

    return true
}

func dfs(g [][]int, status []int, start int) bool {
    status[start] = 1

    for _, nei := range g[start] {
        var ts int = status[nei]
        if ts == 1 {
            return false
        }
        if ts == 0 {
            if !dfs(g, status, nei) {
                return false
            }
        }
    }

    status[start] = 2

    return true
}