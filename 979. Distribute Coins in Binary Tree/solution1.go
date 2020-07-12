type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func distributeCoins(root *TreeNode) int {
    _, res := helper(root)
    return res
}

func helper(root *TreeNode) (int, int) {
    if root == nil {
        return 0, 0
    }

    le, a := helper(root.Left)
    ri, b := helper(root.Right)

    return 1 - root.Val + le + ri, a + b + abs(le) + abs(ri)
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}