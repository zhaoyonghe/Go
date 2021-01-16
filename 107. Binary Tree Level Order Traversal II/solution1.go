type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
    var res [][]int
    lfs(root, 0, &res)
    for i := 0; i < len(res) / 2; i++ {
        res[i], res[len(res) - 1 - i] = res[len(res) - 1 - i], res[i] 
    }
    return res
}

func lfs(root *TreeNode, level int, res *[][]int) {
    if root == nil {
        return
    }
    if len(*res) == level {
        *res = append(*res, nil)
    }
    lfs(root.Left, level + 1, res)
    (*res)[level] = append((*res)[level], root.Val)
    lfs(root.Right, level + 1, res)
}