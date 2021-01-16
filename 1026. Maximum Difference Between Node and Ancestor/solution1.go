/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diff(a int, b int) int {
    if a > b {
        return a - b
    } else {
        return b - a
    }
}

func max(a int, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}

func min(a int, b int) int {
    if a < b {
        return a
    } else {
        return b
    }
}

func maxAncestorDiff(root *TreeNode) int {
    return getMax(root, root.Val, root.Val)
}

func getMax(root *TreeNode, ma int, mi int) int {
    if root == nil {
        return 0
    }

    var temp = max(diff(ma, root.Val), diff(mi, root.Val))
    
    ma = max(ma, root.Val)
    mi = min(mi, root.Val)

    temp = max(temp, getMax(root.Left, ma, mi))
    temp = max(temp, getMax(root.Right, ma, mi))

    return temp
}