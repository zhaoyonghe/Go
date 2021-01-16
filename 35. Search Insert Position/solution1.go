func searchInsert(nums []int, target int) int {
    for i, val := range nums {
        if target <= val {
            return i
        }
    }

    return len(nums)
}