func max(a int, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}

func lengthOfLongestSubstring(s string) int {
    var m = make(map[rune] int)

    var res int = 0
    var base int = -1

    for i, b: = range s {
        if index, ok: = m[b]; ok {
            res = max(res, i - max(index, base))
            base = max(index, base)
        } else {
            res = max(res, i - base)
        }
        m[b] = i
    }

    return res
}