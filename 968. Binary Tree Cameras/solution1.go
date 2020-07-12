type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func minCameraCover(root *TreeNode) int {
    res := dfs(root)
    return min(res[0], res[1])
}

func dfs(root *TreeNode) []int {
    if root == nil {
        return []int{10000000, 0, 0}
    }

    le := dfs(root.Left)
    ri := dfs(root.Right)

    return []int{
        1 + min3(le) + min3(ri),
        min3([]int{le[0] + ri[1], le[1] + ri[0], le[0] + ri[0]}),
        le[1] + ri[1],
    }
}

func min(a, b int) int {
    if a > b {
        return b
    }
    return a
}

func min3(a []int) int {
    return min(min(a[0], a[1]), a[2])
}