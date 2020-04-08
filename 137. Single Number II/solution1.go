func get(arr [64]int) int {
    var res int32 = 0

    for i := 0; i < 64; i++ {
        if arr[i] != 0 {
            res = res | (1 << i)
        }
    }

    return int(res)
}


func singleNumber(nums []int) int {
    var arr [64]int

    for _, val := range nums {
        for i := 0; i < 64; i++ {
            if ((1 << i) & val) != 0 {
                // the i-th bit is 1
                arr[i] = (arr[i] + 1) % 3
            }
        }
    }
    
    return get(arr)
}