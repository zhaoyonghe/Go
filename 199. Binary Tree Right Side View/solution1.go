type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
    var res []int
    helper(root, &res, 0)
    return res
}

func helper(root *TreeNode, res *[]int, level int) {
    if root == nil {
        return
    }
    if len(*res) == level {
        *res = append(*res, root.Val)
    }
    (*res)[level] = root.Val
    helper(root.Left, res, level+1)
    helper(root.Right, res, level+1)
}