type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func flatten(root *TreeNode)  {
    helper(root)
}

func helper(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }

    le := helper(root.Left)
    ri := helper(root.Right)

    root.Left = nil
    root.Right = le
    cur := root
    for cur.Right != nil {
        cur = cur.Right
    }
    cur.Right = ri

    return root
}