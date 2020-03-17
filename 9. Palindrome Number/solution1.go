func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }
    
    var y int64 = (int64)(x)
    var temp int64 = 0
    for y != 0 {
        temp *= 10
        temp += (y % 10)
        y /= 10
        
    }
    
    return temp == (int64)(x)
}