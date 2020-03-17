func min(a int, b int) int {
    if a < b {
        return a
    } else {
        return b
    }
}

func max(a int, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}

func maxArea(height []int) int {
    var i int = 0
    var j int = len(height) - 1
    
    var res int = 0
    
    for i < j {
        res = max(res, (j - i) * min(height[i], height[j]))
        if height[i] < height[j] {
            i++
        } else {
            j--
        }
    }
    
    return res
}