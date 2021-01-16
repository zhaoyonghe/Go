func romanToInt(s string) int {
    var res int = 0

    for i:= 0; i < len(s); i++ {
        switch (s[i]) {
            case 'I':
                if i == len(s) - 1 || (s[i + 1] != 'V' && s[i + 1] != 'X') {
                    res += 1
                } else {
                    res -= 1
                }
            case 'V':
                res += 5
            case 'X':
                if i == len(s) - 1 || (s[i + 1] != 'L' && s[i + 1] != 'C') {
                    res += 10
                } else {
                    res -= 10
                }
            case 'L':
                res += 50
            case 'C':
                if i == len(s) - 1 || (s[i + 1] != 'D' && s[i + 1] != 'M') {
                    res += 100
                } else {
                    res -= 100
                }
            case 'D':
                res += 500
            case 'M':
                res += 1000
        }
    }

    return res
}