func productExceptSelf(nums []int) []int {
    if len(nums) == 1 {
        return nums;
    }

    n := len(nums)

    var result []int = make([]int, n)
    temp := nums[0]
    for i := 1; i < n; i++ {
        result[i] = temp
        temp *= nums[i]
    }
    result[0] = 1
    temp = nums[n - 1]
    for i := n - 2; i >= 0; i-- {
        result[i] *= temp
        temp *= nums[i]
    }

    return result
}