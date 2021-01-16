func min(a int, b int) int {
    if a < b {
        return a
    } else {
        return b
    }
}

func trap(height []int) int {
    if len(height) == 0 {
        return 0
    }

    level := 0
    res := 0
    i := 0
    j := len(height) - 1

    for i < j {
        //fmt.Printf("%v,%v\n", i, j)
        res += (min(height[i], height[j]) - level) * (j - i + 1)

        a := height[i]
        b := height[j]
        level = min(a, b)

        if a > b {
            for i < j && height[j] <= b {
                j--
            }
        } else if a < b {
            for i < j && height[i] <= a {
                i++
            }
        } else {
            for i < j && height[i] <= a {
                i++
            }
            for i < j && height[j] <= b {
                j--
            }
        }  
    }

    res += (height[i] - level)

    for _, val := range height {
        res -= val
    }

    return res
}