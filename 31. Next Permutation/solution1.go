func reverse(nums []int, i int, j int) {
    for k := 0; k <= (j - i - 1) / 2; k++ {
        nums[i + k], nums[j - k] = nums[j - k], nums[i + k]
    }
}

func nextPermutation(nums []int)  {
    var n int = len(nums)
    var i int = n - 2
    for i >= 0 {
        if nums[i] < nums[i + 1] {
            var j int = i + 1
            for j < n && nums[i] < nums[j] {
                j++
            }
            nums[i], nums[j - 1] = nums[j - 1], nums[i]
            reverse(nums, i + 1, n - 1)
            break
        }
        i--
    }
    if i == -1 {
        reverse(nums, 0, n - 1)
    }
}