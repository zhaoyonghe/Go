func singleNumber(nums []int) int {
    var res int = 0

    for _, val := range nums {
        res ^= val
    }

    return res
}