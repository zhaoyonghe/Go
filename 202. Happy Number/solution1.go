func getSum(n int) int {
    res := 0
    for n != 0 {
        temp := n % 10
        res += (temp * temp)
        n /= 10
    }
    return res
}

func isHappy(n int) bool {
    m := make(map[int]bool)
    m[n] = true
    for {
        n = getSum(n)
        if n == 1 {
            return true
        }
        if has, _ := m[n]; has {
            return false
        }
        m[n] = true
    }
}