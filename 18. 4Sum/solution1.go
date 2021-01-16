func fourSum(nums []int, target int) [][]int {
    var res [][]int
    
    sort.Ints(nums)
    
    var i int = 0
    
    for i < len(nums) {
        var j int = i + 1
        
        for j < len(nums) {
            var m int = j + 1
            var n int = len(nums) - 1
            
            var target int = target - nums[i] - nums[j]
            
            for m < n {
                if nums[m] + nums[n] > target {
                    n--
                } else if nums[m] + nums[n] < target {
                    m++
                } else {
                    res = append(res, []int{nums[i], nums[j], nums[m], nums[n]})
                    n--
                    for m < n && nums[n + 1] == nums[n] {
                        n--
                    }
                    m++
                    for m < n && nums[m - 1] == nums[m] {
                        m++
                    }
                }
            }
            
            j++
            for j < len(nums) && nums[j - 1] == nums[j] {
                j++
            }
        }
        
        i++
        for i < len(nums) && nums[i - 1] == nums[i] {
            i++
        }
    }
    
    return res
}