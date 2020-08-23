type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
    if root == nil {
        return 0
    }
    var res int
    helper(root, &res)
    return res - 1
}

func helper(root *TreeNode, res *int) int {
    if root == nil {
        return 0
    }
    lp := helper(root.Left, res)
    rp := helper(root.Right, res)
    *res = max(*res, 1+lp+rp)
    return 1 + max(lp, rp)
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}