type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0 {
        return nil
    }

    node := &TreeNode{
        Val: preorder[0],
    }

    mid := find(inorder, preorder[0])
    node.Left = buildTree(preorder[1:mid+1], inorder[:mid])
    node.Right = buildTree(preorder[mid+1:], inorder[mid+1:])
    return node
}

func find(arr []int, num int) int {
    for i, n := range arr {
        if n == num {
            return i
        }
    }
    return -1
}