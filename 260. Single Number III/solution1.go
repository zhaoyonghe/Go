func singleNumber(nums []int) []int {
    var div int = 0

    for _, val := range nums {
        div ^= val
    }

    for i := 0; i < 64; i++ {
        if ((1 << i) & div) != 0 {
            div = (1 << i) & div
            break
        }
    }

    var res []int = make([]int, 2)
    for _, val := range nums {
        if (div & val) == 0 {
            res[0] ^= val
        } else {
            res[1] ^= val
        }
    }

    return res
}