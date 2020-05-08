func findMinHeightTrees(n int, edges [][]int) []int {
    if n == 1 {
        return []int{0}
    }
    if n == 2 {
        return []int{0, 1}
    }

    for {
        // fill in degrees
        degrees := make([]int, n)
        for _, val := range edges {
            degrees[val[0]]++
            degrees[val[1]]++
        }
        // remove edges with leaves
        temp := [][]int{}
        for _, val := range edges {
            if degrees[val[0]] > 1 && degrees[val[1]] > 1 {
                temp = append(temp, val)
            }
        }
        edges = temp
        if len(edges) == 0 {
            for i, val := range degrees {
                if val > 1 {
                    return []int{i}
                }
            }
        }
        if len(edges) == 1 {
            return edges[0]
        }
    }
    return []int{-1, -1}
}