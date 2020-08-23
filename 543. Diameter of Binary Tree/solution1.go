type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
    if root == nil {
        return 0
    }
    _, res := helper(root)
    return res - 3
}

func helper(root *TreeNode) (int, int) {
    if root == nil {
        return 1, 1
    }

    lh, llp := helper(root.Left)
    rh, rlp := helper(root.Right)
    
    return max(lh, rh) + 1, max(max(llp, rlp), lh + rh + 1)
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}