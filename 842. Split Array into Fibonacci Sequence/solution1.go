func splitIntoFibonacci(S string) []int {
    first := int64(0)
    for i := 0; i < len(S) && i < 10; i++ {
        first *= 10
        first += int64(S[i] - '0')
        if first > math.MaxInt32 { 
            break
        }
        second := int64(0)
        for j := i + 1; j < len(S) && j < i + 11; j++ {
            second *= 10
            second += int64(S[j] - '0')
            if second > math.MaxInt32 { 
                break
            }
            if can, res := check(S, first, second, j + 1); can {
                return res
            }
            if second == 0 {
                break
            }
        }
        if first == 0 {
            break
        }
    }
    return []int{}
}

func check(s string, first, second int64, i int) (bool, []int) {
    res := []int{int(first), int(second)}
    for {
        first += second
        if first > math.MaxInt32 {
            return false, nil
        }
        next := strconv.Itoa(int(first))
        if !strings.HasPrefix(s[i:], next) {
            return false, nil
        }
        res = append(res, int(first))
        i += len(next)
        if i == len(s) {
            return true, res
        }
        first, second = second, first
    }
    return false, nil
}