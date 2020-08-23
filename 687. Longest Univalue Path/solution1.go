type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func longestUnivaluePath(root *TreeNode) int {
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

    up := 1

    if root.Left != nil && root.Left.Val == root.Val {
        up += lp
    } else {
        lp = 0
    }
    if root.Right != nil && root.Right.Val == root.Val  {
        up += rp
    } else {
        rp = 0
    }

    *res = max(*res, up)
    return max(lp, rp) + 1
} 

func max(a, b int) int {
    if a > b {
        return a
    } 
    return b
}