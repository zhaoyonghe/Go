func strStr(haystack string, needle string) int {
    if haystack == needle || needle == "" {
        return 0
    }
    
    if len(haystack) < len(needle) {
        return -1
    }
    
    for i, _ := range haystack[:len(haystack) - len(needle) + 1] {
        if haystack[i: i + len(needle)] == needle {
            return i
        }
    }
    
    return -1
}