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

func findMedianSortedArrays(nums1[] int, nums2[] int) float64 {
    var m int = len(nums1)
    var n int = len(nums2)

    if m < n {
        return findMedianSortedArrays(nums2, nums1)
    }

    if n == 0 {
        if (m + n) % 2 == 0 {
            return float64(nums1[(m + n) / 2 - 1] + nums1[(m + n) / 2]) / 2.0
        } else {
            return float64(nums1[(m + n) / 2])
        }
    }

    var bl int = -1
    var al int = (m + n + 1) / 2 - bl - 2

    var low int = -1
    var high int = n

    for {
        //fmt.Printf("%v, %v\n", al, bl)
        var lefta int = math.MinInt32
        var leftb int = math.MinInt32
        var righta int = math.MaxInt32
        var rightb int = math.MaxInt32
        if al >= 0 && al < m {
            lefta = nums1[al]
        }
        if bl >= 0 && bl < n {
            leftb = nums2[bl]
        }
        if al + 1 >= 0 && al + 1 < m {
            righta = nums1[al + 1]
        }
        if bl + 1 >= 0 && bl + 1 < n {
            rightb = nums2[bl + 1]
        }
        if lefta <= rightb && leftb <= righta {
            if (m + n) % 2 == 0 {
                return float64(max(lefta, leftb) + min(righta, rightb)) / 2.0
            } else {
                return float64(max(lefta, leftb))
            }
        } else if lefta > rightb {
            low = bl + 1
            bl = (low + high) / 2
            al = (m + n + 1) / 2 - bl - 2
        } else {
            // leftb > righta
            high = bl - 1
            bl = (low + high) / 2
            al = (m + n + 1) / 2 - bl - 2
        }
    }

    return 0.0
}