type Pair struct {
    num int
    count int
}

func combinationSum(candidates []int, target int) [][]int {
    var m map[int]int = make(map[int]int)
    for _, val := range candidates {
        m[val] = target / val
    }

    var can []Pair = []Pair{}
    for k, v := range m {
        can = append(can, Pair{k, v})
    }

    var res [][]int = [][]int{}
    getCombination(can, target, &[]int{}, &res)
    return res
}

func getCombination(candidates []Pair, target int, cur *[]int, res *[][]int) {
    if target == 0 {
        var temp []int = []int{}
        for _, val := range *cur {
            temp = append(temp, val)
        }
        *res = append(*res, temp)
        return
    }

    if target < 0 || len(candidates) == 0 {
        return
    }

    getCombination(candidates[1:], target, cur, res)
    
    for i := 0; i < candidates[0].count; i++ {
        *cur = append(*cur, candidates[0].num)
        getCombination(candidates[1:], target - (i + 1) * candidates[0].num, cur, res)
    }
    *cur = (*cur)[:len(*cur) - candidates[0].count]

    return
}