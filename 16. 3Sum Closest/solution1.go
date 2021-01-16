func abs(a int) int {
    if a > 0 {
        return a
    } else {
        return -a
    }
}

func threeSumClosest(nums []int, target int) int {
    var offset int = math.MaxInt32
    
    sort.Ints(nums)
    
    var i int = 0
    
    for i < len(nums) {
        var aim int = target - nums[i]
        
        var j int = i + 1
        var k int = len(nums) - 1
        
        for j < k {
            var off int = nums[i] + nums[j] + nums[k] - target
            if off == 0 {
                return target
            }
            if abs(off) < abs(offset) {
                offset = off
            }
            
            if nums[j] + nums[k] > aim {
                k--
            } else if nums[j] + nums[k] < aim {
                j++
            } else {
                j++
                k--
            }
        }
        
        i++
    }
    
    return target + offset
}