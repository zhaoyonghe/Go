func diffWaysToCompute(input string) []int {
    var ns []int
    var ops []byte
    num := 0
    for i := 0; i < len(input); i++ {
        if '0' <= input[i] && input[i] <= '9' {
            num *= 10
            num += int(input[i] - '0')
        } else {
            // operator
            ns = append(ns, num)
            ops = append(ops, input[i])
            num = 0
        }
    }
    ns = append(ns, num)
    return helper(ns, ops, 0, len(ops), make(map[[2]int][]int))
}

func helper(ns []int, ops []byte, s, e int, memo map[[2]int][]int) []int {
    if s == e {
        return []int{ns[s]}
    }
    if val, ok := memo[[2]int{s,e}]; ok {
        return val
    }
    
    var res []int
    for i := s; i < e; i++ {
        le := helper(ns, ops, s, i, memo)
        ri := helper(ns, ops, i+1, e, memo)
        op := ops[i]
        for _, lval := range le {
            for _, rval := range ri {
                res = append(res, calculate(lval, rval, op))
            }
        } 
    }
    memo[[2]int{s,e}] = res
    return res
}

func calculate(a, b int, op byte) int {
    switch op {
        case '+':
        return a + b
        case '-':
        return a - b
        case '*':
        return a * b
        case '/':
        return a / b
    }
    return 0
}