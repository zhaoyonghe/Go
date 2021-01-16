type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
    if n == 0 {
        return nil
    }
    return helper(1, n)
}

func helper(start, end int) []*TreeNode {
    if start > end {
        return []*TreeNode{nil}
    }

    var allTree []*TreeNode
    for i := start; i <= end; i++ {
        left := helper(start, i - 1)
        right := helper(i + 1, end)
        for _, le := range left {
            for _, ri := range right {
                allTree = append(allTree, &TreeNode{i,le,ri})
            }
        }
    }
    return allTree
}