func min(a, b int) int {
    if a > b {
        return b
    } else {
        return a
    }
}

func findCheapestPrice(n int, flights [][]int, src int, dst int, K int) int {
    distance := make([]int, n)
    for i := 0; i < n; i++ {
        distance[i] = -1
    }
    distance[src] = 0
    
    for i := 0; i <= K; i++ {
        tmp := make([]int, n)
        copy(tmp, distance)
        for _, val := range flights {
            u := val[0]
            v := val[1]
            w := val[2]

            if distance[u] != -1 {
                if tmp[v] == -1 {
                    tmp[v] = distance[u] + w
                } else {
                    tmp[v] = min(tmp[v], distance[u] + w)
                }
            } 
        }
        distance = tmp
        //fmt.Printf("%v\n", distance)
        //fmt.Printf("%v\n", K)
    }

    return distance[dst]
}