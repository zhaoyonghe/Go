func twoSum(nums[] int, target int)[] int {
    var m = make(map[int] int)
    for i, num: = range nums {
        if val, ok: = m[target - num];
        ok {
            return [] int {
                i, val
            }
        }
        m[num] = i
    }
    return [] int {
        0, 0
    }
}