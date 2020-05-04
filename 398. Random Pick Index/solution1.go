type Solution struct {
    nums []int
}


func Constructor(nums []int) Solution {
    return Solution{nums:nums}
}


func (this *Solution) Pick(target int) int {
    var count int = 1
    var res int = -1
    for i, val := range this.nums {
        if val == target {
            if rand.Intn(count) == 0 {
                res = i
            }
            count++
        }
    }
    return res
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */