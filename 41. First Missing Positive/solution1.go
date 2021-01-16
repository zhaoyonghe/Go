func firstMissingPositive(nums []int) int {
    i := 0
    for i < len(nums) {
        for 1 <= nums[i] && nums[i] <= len(nums) && nums[i] != nums[nums[i] - 1] {
            nums[i], nums[nums[i] - 1] = nums[nums[i] - 1], nums[i]
        } 
        i++
    }

    for i, val := range nums {
        if val != i + 1 {
            return i + 1
        }
    }

    return len(nums) + 1
}