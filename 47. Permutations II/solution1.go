func permuteUnique(nums []int) [][]int {
    var res [][]int
    permute(nums, 0, &res)
    return res
}

func permute(nums []int, start int, res *[][]int) {
    if start == len(nums) {
        var temp []int = make([]int, len(nums))
        copy(temp, nums)
        *res = append(*res, temp)
        return
    }
loop:
    for i := start; i < len(nums); i++ {
        for j := start; j < i; j++ {
            if nums[i] == nums[j] {
                continue loop
            }
        }
        nums[start], nums[i] = nums[i], nums[start]
        permute(nums, start + 1, res)
        nums[start], nums[i] = nums[i], nums[start]
    }
}