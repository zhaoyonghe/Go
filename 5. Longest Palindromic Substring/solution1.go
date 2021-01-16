func max(a int, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}

func longestPalindrome(s string) string {
    if len(s) == 0 {
        return ""
    }
    
    var res int = 1
    var temp int
    var tempi int
    var tempj int
    var ii int = 0
    var jj int = 0
    
    for i := 0; i < len(s); i++ {
        temp, tempi, tempj = expand(s, i, i)
        if res < temp {
            res = temp
            ii = tempi
            jj = tempj
        }
    }
    
    for i := 0; i < len(s) - 1; i++ {
        temp, tempi, tempj = expand(s, i, i + 1)
        if res < temp {
            res = temp
            ii = tempi
            jj = tempj
        }
    }
    
    // fmt.Printf("%v, %v\n", ii, jj)
    
    return s[ii: jj + 1]
}

func expand(s string, pi int, pj int) (int, int, int) {
    var i int = pi
    var j int = pj
    var length int = 0
    
    if pi == pj {
        length = -1    
    }
    
    for i >= 0 && j < len(s) {
        if s[i] == s[j] {
            length += 2
        } else {
            break
        }
        i -= 1
        j += 1
    }
    
    if j - i + 1 > length {
        i += 1
        j -= 1
    }
    
    return length, i, j
}