func longestCommonPrefix(strs[] string) string {
    if len(strs) == 0 {
        return ""
    }

    var c byte
    var j int = 0
    var res bytes.Buffer

    out:
        for {
            //fmt.Printf("%v\n", j)
            if j == len(strs[0]) {
                break
            } else {
                c = strs[0][j]
            }

            for _, s: = range strs {
                if j == len(s) || s[j] != c {
                    break out
                }
            }
            res.WriteByte(c)
            j++
        }

    return res.String()
}