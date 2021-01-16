func max(a int, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}

func min(a int, b int) int {
    if a < b {
        return a
    } else {
        return b
    }
}

func findKthLargest(nums []int, k int) int {
    var n int = len(nums)

    if n == 1 {
        return nums[0]
    }

    var r int = rand.Intn(n)
    nums[r], nums[n - 1] = nums[n - 1], nums[r]

    var i int = -1
    var j int = n - 2

    for i != j {
        if nums[i + 1] <= nums[n - 1] {
            i++
        } else {
            nums[i + 1], nums[j] = nums[j], nums[i + 1]
            j--
        }
    }
    // i == j
    nums[i + 1], nums[n - 1] = nums[n - 1], nums[i + 1]
    i++
    
    //fmt.Printf("%v, %v\n", nums, i)

    if n - i > k {
        return findKthLargest(nums[i+1:], k)
    } else if  n - i < k {
        return findKthLargest(nums[:i], k - n + i)
    } else {
        //  n - i == k
        return nums[i]
    }
}