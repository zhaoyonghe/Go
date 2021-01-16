func reverse(arr *[]byte) *[]byte {
    length := len(*arr)
    for i := 0; i < length / 2; i++ {
        (*arr)[i], (*arr)[length - 1 - i] = (*arr)[length - 1 - i], (*arr)[i]
    }
    return arr
}

func addStrings(num1 string, num2 string) string {
    var res bytes.Buffer

    var i int = len(num1) - 1
    var j int = len(num2) - 1
    var carry int = 0

    for i >= 0 || j >= 0 {
        var a int = 0
        var b int = 0
        if i >= 0 {
            a = int(num1[i] - '0')
        }
        if j >= 0 {
            b = int(num2[j] - '0')
        }
        temp := a + b + carry
        res.WriteByte(byte(temp % 10 + int('0')))
        carry = temp / 10
        i--
        j--
    }

    if carry == 1 {
        res.WriteByte('1')
    }

    var temp []byte = res.Bytes()

    return bytes.NewBuffer(*reverse(&temp)).String()
}