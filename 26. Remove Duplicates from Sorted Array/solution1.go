func removeDuplicates(nums []int) int {
    if len(nums) <= 1 {
        return len(nums)
    }
    
    var j int = 0
    
    for _, val := range nums[1:] {
        if val != nums[j] {
            nums[j + 1] = val
            j++
        }
    }
    
    return j + 1
}