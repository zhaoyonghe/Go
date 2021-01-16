func search(nums []int, target int) int {
    if nums == nil || len(nums) == 0 {
        return -1
    }

    var low int = 0
    var high int = len(nums) - 1

    for low <= high {
        var mid int = (low + high) / 2
        if target == nums[mid] {
            return mid
        }
        if nums[low] < nums[mid] {
            if nums[low] <= target && target < nums[mid] {
                high = mid - 1
            } else {
                low = mid + 1
            }
        } else if nums[low] > nums[mid] {
            if nums[mid] < target && target <= nums[high] {
                low = mid + 1
            } else {
                high = mid - 1
            }
        } else {
            // nums[low] == nums[mid] != target
            // low == mid
            low += 1
        }
        // fmt.Printf("%v\n", mid)
    }

    return -1
}