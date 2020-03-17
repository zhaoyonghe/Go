func reverse(x int) int {
    var res int64 = 0
    for x != 0 {
        res *= 10
        res += (int64)(x % 10)
        x /= 10
    }
    
    if res > math.MaxInt32 || res < math.MinInt32 {
        return 0
    }
    
    return (int)(res)
}