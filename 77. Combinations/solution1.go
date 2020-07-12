func combine(n int, k int) [][]int {
    res := [][]int{}
    if n < k || n < 1 || k < 1 {
        return res
    }

    model := make([]int, n)
    for i := 1; i <= n; i++ {
        model[i - 1] = i
    }
    
    helper(model, 0, k, &[]int{}, &res)

    return res
}

func helper(model []int, s int, k int, cur *[]int, res *[][]int) {
    // fmt.Printf("%v,%v\n", s, k)
    if len(model[s:]) < k {
        return
    }
    if k == 0 {
        tmp := []int{}
        tmp = append(tmp, (*cur)...)
        *res = append(*res, tmp)
        return        
    }
    if len(model[s:]) == k {
        tmp := []int{}
        tmp = append(tmp, (*cur)...)
        tmp = append(tmp, model[s:]...)
        *res = append(*res, tmp)
        return        
    }
    // select s
    se := append(*cur, model[s])
    helper(model, s + 1, k - 1, &se, res)
    // no select s
    helper(model, s + 1, k, cur, res)
    return 
}