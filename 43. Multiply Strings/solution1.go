func add(num1 *[]int, num2 []int, num2pow int) {
    var i int = 0
    var j int = -num2pow
    var carry int = 0

    for i < len(*num1) || j < len(num2) {
        var a int = 0
        var b int = 0
        if i < len(*num1) {
            a = (*num1)[i]
        }
        if 0 <= j && j < len(num2) {
            b = num2[j]
        }
        temp := a + b + carry
        if i < len(*num1) {
            (*num1)[i] = temp % 10
        } else {
            *num1 = append(*num1, temp % 10)
        }
        
        carry = temp / 10
        i++
        j++
    }

    if carry == 1 {
        *num1 = append(*num1, carry)
    }
}


func mult(num []int, a int, res *[]int) {
    var carry int = 0
    for i, _ := range num {
        temp := carry + num[i] * a
        *res = append(*res, temp % 10)
        carry = temp / 10
    }

    if carry > 0 {
        *res = append(*res, carry)
    }
}

func multiply(num1 string, num2 string) string {
    if num1 == "0" || num2 == "0" {
        return "0"
    }

    var n1 []int
    var n2 []int

    for i := len(num1) - 1; i >= 0; i-- {
        n1 = append(n1, int(num1[i] - '0'))
    }
    for j := len(num2) - 1; j >= 0; j-- {
        n2 = append(n2, int(num2[j] - '0'))
    }

    var cache [10][]int
    cache[0] = []int{0}
    cache[1] = n1

    var res []int = []int{0}

    for j, a := range n2 {
        if len(cache[a]) == 0 {
            mult(n1, a, &(cache[a]))
        }

        add(&res, cache[a], j)
    }

    var ans bytes.Buffer
    for i := len(res) - 1; i >= 0; i-- {
        ans.WriteByte(byte(res[i] + int('0')))
    }

    return ans.String()
}