func subsets(nums []int) [][]int {
    res := [][]int{}
    helper(nums, 0, &[]int{}, &res)
    return res
}

func helper(nums []int, i int, cur *[]int, res *[][]int) {
    if i == len(nums) {
        tmp := append([]int{}, (*cur)...)
        *res = append(*res, tmp)
        return
    }
    // select i
    *cur = append(*cur, nums[i])
    helper(nums, i + 1, cur, res)
    *cur = pop(*cur)
    // no select i
    helper(nums, i + 1, cur, res)
}

func pop(a []int) []int {
    return a[:len(a) - 1]
} 