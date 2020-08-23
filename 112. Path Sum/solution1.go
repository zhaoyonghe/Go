type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func hasPathSum(root *TreeNode, sum int) bool {
    if root == nil {
        return false
    }

    return helper(root, sum)
}

func helper(root *TreeNode, sum int) bool {
    if root.Left == nil && root.Right == nil {
        return sum == root.Val
    }

    if root.Left != nil && helper(root.Left, sum - root.Val) {
        return true 
    }
    if root.Right != nil && helper(root.Right, sum - root.Val) {
        return true 
    }

    return false
}