func combinationSum3(k int, n int) [][]int {
    res := &[][]int{}
    helper(k, 0, n, &[]int{}, res)
    return *res
}

func helper(k, bgt, n int, cur *[]int, res *[][]int) {
    if k == 0 {
        if n == 0 {
            tmp := []int{}
            tmp = append(tmp, (*cur)...)
            *res = append(*res, tmp)
        }
        return
    }

    for i := bgt + 1; i <= 9; i++ {
        *cur = append(*cur, i)
        helper(k - 1, i, n - i, cur, res)
        *cur = (*cur)[:len(*cur) - 1]
    }
}