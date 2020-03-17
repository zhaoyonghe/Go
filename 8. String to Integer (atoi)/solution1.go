func isDigit(c byte) bool {
    return c >= '0' && c <= '9'
}

func getNum(s string, i int, sign int64) int {
    var res int64 = 0
    var count int = 0
    
    // filter leading zero
    for i < len(s) && s[i] == '0' {
        i++
    }
    
    for i < len(s) && count < 11 {
        if isDigit(s[i]) {
            res *= 10
            res += (int64)(s[i] - '0')
            count++
            i++
        } else {
            break
        }
    }
    
    res *= sign
    
    if res > math.MaxInt32 {
        return math.MaxInt32
    } else if res < math.MinInt32 {
        return math.MinInt32
    } else {
        return (int)(res)
    }
}

func myAtoi(str string) int {
    
    var i = 0
    
    for i < len(str) {
        if str[i] == ' ' {
            i++
        } else if !isDigit(str[i]) && str[i] != '+' && str[i] != '-' {
            return 0
        } else {
            if str[i] == '+' {
                return getNum(str, i + 1, 1)
            } else if str[i] == '-' {
                return getNum(str, i + 1, -1)
            } else {
                // str[i] is digit
                return getNum(str, i, 1)
            }
        }
        
    }
    
    return 0
}