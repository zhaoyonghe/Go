func searchInsert(nums []int, target int) int {
    if target > nums[len(nums) - 1] {
        return len(nums)
    }
    
    var low int = 0
    var high int = len(nums) - 1

    for low < high {
        var mid int = (low + high) / 2

        if target <= nums[mid] {
            high = mid
        } else {
            low = mid + 1
        }
    }

    return low
}