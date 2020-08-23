func fixedPoint(A []int) int {
    if len(A) == 0 {
        return -1
    }
    low, high := 0, len(A) - 1
    for low < high {
        mid := (low + high) / 2
        if A[mid] < mid {
            low = mid + 1
        } else if A[mid] > mid {
            high = mid - 1
        } else {
            high = mid
        }
    }
    if A[low] == low {
        return low
    }
    return -1
}