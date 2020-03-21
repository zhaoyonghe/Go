func isValid(s string) bool {
    var st []byte
    
    for _, c := range s {
        switch(c) {
            case '(':
            st = append(st, '(')
            case '{':
            st = append(st, '{')
            case '[':
            st = append(st, '[')
            case ')':
            if len(st) == 0 || st[len(st) - 1] != '(' {
                return false
            } else {
                st = st[0:len(st) - 1]
            }
            case '}':
            if len(st) == 0 || st[len(st) - 1] != '{' {
                return false
            } else {
                st = st[0:len(st) - 1]
            }
            case ']':
            if len(st) == 0 || st[len(st) - 1] != '[' {
                return false
            } else {
                st = st[0:len(st) - 1]
            } 
        }
    }
    
    return len(st) == 0
}