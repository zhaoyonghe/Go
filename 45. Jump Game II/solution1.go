func min(a int, b int) int {
    if a < b {
        return a
    } else {
        return b
    }
}

func fill(arr []int, start int, end int, val int) {
    for i := start; i <= min(end, len(arr) - 1); i++ {
        arr[i] = val
    }
}

func jump(nums []int) int {
    var n int = len(nums)
    if n == 1 {
        return 0
    }

    var jumps []int = make([]int, n)

    var farest = 0

    for i := 0; i < n; i++ {
        if i + nums[i] > farest {
            fill(jumps, farest + 1, i + nums[i], jumps[i] + 1)
            farest = i + nums[i]
        }
        //fmt.Printf("%v\n", jumps)
        if farest >= n - 1 {
            return jumps[n - 1]
        }
    }

    return -1
}