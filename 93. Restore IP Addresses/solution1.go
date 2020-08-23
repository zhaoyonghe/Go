func restoreIpAddresses(s string) []string {
    var res []string
    restore(s, 0, &[]string{}, &res)
    return res
}

func valid(n string) bool {
    if n[0] == '0' && len(n) > 1 {
        return false
    }
    num, err := strconv.Atoi(n)
    return err == nil && num <= 255 && num >= 0
}

func restore(s string, i int, cur *[]string, res *[]string) {
    if len(*cur) == 4 && i < len(s) {
        return
    }
    if i == len(s) {
        if len(*cur) == 4 {
            *res = append(*res, strings.Join(*cur, "."))
        }
        return
    }
    for j := i + 1; j <= len(s) && j <= i + 3; j++ {
        tmp := s[i:j]
        if valid(tmp) {
            *cur = append(*cur, tmp)
            restore(s, j, cur, res)
            *cur = (*cur)[0:len(*cur) - 1]
        }
    } 
}