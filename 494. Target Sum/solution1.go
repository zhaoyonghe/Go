func findTargetSumWays(nums []int, S int) int {
    m := make(map[[2]int]int)
    return dfs(nums, S, 0, m)
}

func dfs(nums []int, sum int, i int, m map[[2]int]int) int {
    if i == len(nums) {
        if sum == 0 {
            return 1
        } else {
            return 0
        }
    }

    res := 0

    if val, ok := m[[2]int{sum + nums[i], i + 1}]; ok {
        res += val
    } else {
        res += dfs(nums, sum + nums[i], i + 1, m)
    }

    if val, ok := m[[2]int{sum - nums[i], i + 1}]; ok {
        res += val
    } else {
        res += dfs(nums, sum - nums[i], i + 1, m)
    }

    m[[2]int{sum, i}] = res

    return res
}