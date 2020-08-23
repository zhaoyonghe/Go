func partition(s string) [][]string {
    var res [][]string
    helper(s, 0, &[]string{}, &res)
    return res
}

func helper(str string, low int, cur *[]string, res *[][]string) {
    if low == len(str) {
        var tmp []string
        tmp = append(tmp, (*cur)...)
        *res = append(*res, tmp)
        return
    }
    for i := low; i < len(str); i++ {
        if isPal(str, low, i) {
            *cur = append(*cur, str[low:i+1])
            helper(str, i+1, cur, res)
            *cur = (*cur)[:len(*cur)-1]
        }
    } 
}

func isPal(s string, i, j int) bool {
    for i < j {
        if s[i] != s[j] {
            return false
        }
        i++
        j--
    }
    return true
}