func canPartitionKSubsets(nums []int, k int) bool {
    if len(nums) < k {
        return false
    }
    sum := 0
    for _, val := range nums {
        sum += val
    }
    if sum % k != 0 {
        return false
    }

    sort.Sort(sort.Reverse(sort.IntSlice(nums)))

    target := sum / k
    if nums[0] > target {
        return false
    }

    bucket := make([]int, k)
    for i := 0; i < k; i++ {
        bucket[i] = target
    }

    return helper(nums, bucket, 0, k)
}

func helper(nums []int, bucket []int, i int, k int) bool {
    if i == len(nums) {
        return true
    }
    for j := 0; j < k; j++ {
        if j > 0 && bucket[j] == bucket[j - 1] {
            continue
        }
        if bucket[j] >= nums[i] {
            bucket[j] -= nums[i]
            if helper(nums, bucket, i + 1, k) {
                return true
            }
            bucket[j] += nums[i]
        }
    }
    return false
}