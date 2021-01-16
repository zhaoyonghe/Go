func permute(nums []int) [][]int {
    var res [][]int = [][]int{}

    getPermute(nums, 0, &res)
    return res
}

func getPermute(cur []int, i int, res *[][]int) {
    if i == len(cur) {
        var temp []int
        for _, val := range cur {
            temp = append(temp, val)
        }
        *res = append(*res, temp)
        return
    }

    for j := i; j < len(cur); j++ {
        cur[i], cur[j] = cur[j], cur[i]
        getPermute(cur, i + 1, res)
        cur[i], cur[j] = cur[j], cur[i]
    }
}