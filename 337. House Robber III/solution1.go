type TreeNode struct {
    Val int
    Left *ListNode
    Right *ListNode
}

func rob(root *TreeNode) int {
    return maxes(helper(root))
}

func helper(root *TreeNode) []int {
    if root == nil {
        return []int{0, 0}
    }

    le := helper(root.Left)
    ri := helper(root.Right)

    res := []int{0, 0}
    res[0] = root.Val + le[1] + ri[1]
    res[1] = maxes([]int{le[0] + ri[0], le[0] + ri[1], le[1] + ri[0], le[1] + ri[1]})

    return res
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func maxes(a []int) int {
    res := a[0]
    for _, val := range a {
        res = max(res, val)
    }
    return res
}