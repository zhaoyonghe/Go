type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func pruneTree(root *TreeNode) *TreeNode {
    helper(root)
    return root
}

func helper(root *TreeNode) bool {
    if root == nil {
        return true
    }

    laz := helper(root.Left)
    raz := helper(root.Right)

    if laz {
        root.Left = nil
    }
    if raz {
        root.Right = nil
    }

    return laz && raz && (root.Val == 0)
}