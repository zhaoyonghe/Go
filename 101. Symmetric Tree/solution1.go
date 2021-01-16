type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
    return isSame(root, revert(root))
}

func revert(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    return &TreeNode{
        root.Val,
        revert(root.Right),
        revert(root.Left),
    }
}

func isSame(r1, r2 *TreeNode) bool {
    return r1 == nil && r2 == nil || (r1 != nil && r2 != nil && r1.Val == r2.Val && isSame(r1.Left, r2.Left) && isSame(r1.Right, r2.Right))
}