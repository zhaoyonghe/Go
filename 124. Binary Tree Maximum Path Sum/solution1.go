type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
    res := math.MinInt32
    helper(root, &res)
    return res
}

func helper(root *TreeNode, res *int) int {
    if root == nil {
        return 0
    }
    le := helper(root.Left, res)
    ri := helper(root.Right, res)
    *res = max(*res, root.Val+le+ri)
    return max(max(le+root.Val, ri+root.Val), 0)
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}