func generate(le int, ri int, s string, res *[]string) {
    if le == 0 && ri == 0 {
        *res = append(*res, s)
    }
    
    if le > 0 {
        generate(le - 1, ri, s + "(", res)
    }
    
    if le < ri {
        generate(le, ri - 1, s + ")", res)
    }
}

func generateParenthesis(n int) []string {
    var res []string
    generate(n, n, "", &res)
    return res
}