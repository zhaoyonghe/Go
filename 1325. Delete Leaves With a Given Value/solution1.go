type TreeNode struct {
    Val int
    Left *ListNode
    Right *ListNode
}

func removeLeafNodes(root *TreeNode, target int) *TreeNode {
    if helper(root, target) {
        return nil
    }

    return root
}

func helper(root *TreeNode, target int) bool {
    if root == nil {
        return true
    }

    dle := helper(root.Left, target)
    dri := helper(root.Right, target)

    if dle && dri {
        if root.Val == target {
            return true
        }
    }

    if dle {
        root.Left = nil
    }

    if dri {
        root.Right = nil
    }

    return false
}