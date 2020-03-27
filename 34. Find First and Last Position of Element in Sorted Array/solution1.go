func searchRange(nums []int, target int) []int {
    var res []int = []int{-1, -1}
    if len(nums) == 0 {
        return res
    }
    
    var low int = 0
    var high int = len(nums) - 1

    

    // find first
    for low < high {
        var mid int = (low + high) / 2

        if nums[mid] > target {
            high = mid - 1
        } else if nums[mid] < target {
            low = mid + 1
        } else {
            high = mid
        }
    }

    if nums[low] != target {
        return res
    } else {
        res[0] = low
    }

    low, high = 0, len(nums) - 1

    // find last
    for low < high {
        var mid int = (low + high + 1) / 2

        if nums[mid] > target {
            high = mid - 1
        } else if nums[mid] < target {
            low = mid + 1
        } else {
            low = mid
        }
    }

    res[1] = low

    return res
}