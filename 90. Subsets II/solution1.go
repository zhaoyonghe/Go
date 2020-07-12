type rec struct {
    num int
    times int
}

func subsetsWithDup(nums []int) [][]int {
    sort.Ints(nums)
    res := [][]int{}
    helper(convert(nums), 0, &[]int{}, &res)
    return res
}

func convert(nums []int) []rec {
    res := []rec{}
    for _, val := range nums {
        if len(res) == 0 || res[len(res) - 1].num != val {
            res = append(res, rec{val, 1})
            continue
        }
        res[len(res) - 1].times++
    }
    return res
}

func helper(recs []rec, s int, cur *[]int, res *[][]int) {
    if s == len(recs) {
        tmp := append([]int{}, (*cur)...)
        *res = append(*res, tmp)
        return
    }

    helper(recs, s + 1, cur, res)
    for i := 1; i <= recs[s].times; i++ {
        *cur = append(*cur, recs[s].num)
        helper(recs, s + 1, cur, res)
    }
    *cur = (*cur)[:len(*cur) - recs[s].times]
}