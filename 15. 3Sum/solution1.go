func threeSum(nums []int) [][]int {
    var res [][]int
    
    sort.Ints(nums)
    
    var i int = 0
    
    for i < len(nums) {
        var target int = -nums[i]
        
        var j int = i + 1
        var k int = len(nums) - 1
        
        for j < k {
            if nums[j] + nums[k] > target {
                k--
            } else if nums[j] + nums[k] < target {
                j++
            } else {
                res = append(res, []int{nums[i], nums[j], nums[k]})
                j++
                for j < k && nums[j - 1] == nums[j] {
                    j++
                }
                k--
                for j < k && nums[k + 1] == nums[k] {
                    k--
                }
            }
        }
        
        i++
        for i < len(nums) && nums[i - 1] == nums[i] {
            i++
        }
    }
    
    return res
}