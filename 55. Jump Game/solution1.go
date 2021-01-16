func canJump(nums []int) bool {
    var farest int = 0
    var i int = 0

    for i <= farest {
        if i + nums[i] > farest {
            farest = i + nums[i]
        }
        if farest >= len(nums) - 1 {
                return true
        }
        i++
    }

    return false
}