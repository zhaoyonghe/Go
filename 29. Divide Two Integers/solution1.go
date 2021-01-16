func abs(a int) int {
    if a < 0 {
        return -a
    } else {
        return a
    }
}

func divide(dividend int, divisor int) int {
    if dividend == 0 {
        return 0
    }
    
    if dividend == math.MinInt32 && divisor == -1 {
        return math.MaxInt32
    }
    
    var neg bool = (dividend ^ divisor) < 0
    
    dividend = abs(dividend)
    divisor = abs(divisor)
    
    var result int = 0
    
    for i := 31; i >= 0; i-- {
        if (dividend >> i) >= divisor {
            result += (1 << i)
            dividend -= (divisor << i)
        }
    }
    
    if neg {
        return -result
    } else {
        return result
    }
}