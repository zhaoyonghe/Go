var tab [10]string = [10]string {
    "",
    "",
    "abc",
    "def",
    "ghi",
    "jkl",
    "mno",
    "pqrs",
    "tuv",
    "wxyz" }

func letterCombinations(digits string) []string {
    if len(digits) == 0 {
        return []string{}
    }
    return getCombinations(digits, 0)
}

func getCombinations(digits string, i int) []string {
    if i == len(digits) {
        return []string{""}
    }
    
    var last []string = getCombinations(digits, i + 1)
    var res []string
    
    var temp string = tab[(int)(digits[i] - '0')]
    for _, c := range temp {
        for _, s := range last {
            res = append(res, string(c) + s)
        }
    }
    
    return res
}