func findSubstring(s string, words []string) []int {
    var res []int
    var length = len(words)
    if length == 0 {
        return res
    }
    
    var slen = len(words[0])
    
    var dict = make(map[string]int)
    for _, str := range words {
        val, ok := dict[str]
        if ok {
            dict[str] = val + 1
        } else {
            dict[str] = 1
        }
    }
    
    loop:
    for i := 0; i <= len(s) - slen * length; i++ {
        var j int = i
        var m = make(map[string]int)
        
        for j < i + slen * length {
            var temp string = s[j: j + slen]
            val, ok := dict[temp]
            if !ok {
                continue loop
            }
            val2, ok2 := m[temp]
            if !ok2 {
                m[temp] = 1
            } else {
                if val2 == val {
                    continue loop
                }
                m[temp] = val2 + 1
            }
            
            j += slen
        }
        
        res = append(res, i)
    }
    
    return res
}