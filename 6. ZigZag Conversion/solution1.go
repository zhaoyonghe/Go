func convert(s string, numRows int) string {
    if numRows == 1 || len(s) <= numRows {
        return s
    }
    
    var bb = make([]bytes.Buffer, numRows)
    
    for i := 0; i < numRows; i++ {
        var j int = i
        bb[i].WriteByte(s[j])
        var a int = (numRows - i - 1) * 2 
        var b int = i * 2

        for {

            if j + a < len(s) {
                j += a
                if a != 0 {
                    bb[i].WriteByte(s[j])
                }
            } else {
                break
            }

            if j + b < len(s) {
                j += b
                if b != 0 {
                    bb[i].WriteByte(s[j])
                }
            } else {
                break
            }
        }                
    }
    
    var res string = ""
    for i := 0; i < numRows; i++ {
        //fmt.Printf("%v\n", bb[i].String())
        res += bb[i].String()
    }    
    
    return res
}