func countAndSay(n int) string {
    if n == 1 {
        return "1"
    }

    var s string = countAndSay(n - 1)

    var c byte = s[0]
    var i int = 0
    var count int = 0
    var res string = ""

    for i < len(s) {
        if s[i] == c {
            count++
            i++
        } else {
            res += (strconv.Itoa(count) + (string)(c))
            c = s[i]
            count = 0
        }
    }
    res += (strconv.Itoa(count) + (string)(c))

    return res
}