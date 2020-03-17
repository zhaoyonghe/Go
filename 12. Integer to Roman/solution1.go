func intToRoman(num int) string {
    var tab = [4][10]string{
        {"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"},
        {"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"},
        {"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"},
        {"", "M", "MM", "MMM", "", "", "", "", "", ""}}

    var res string = ""
    for i := 0; i < 4; i++ {
        res = tab[i][num % 10] + res
        num /= 10
    }

    return res
}