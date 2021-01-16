type TreeNode struct {
    Val int
    Left *ListNode
    Right *ListNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
     var pp []*TreeNode
     var qp []*TreeNode

     if !getPath(root, p, &pp) || !getPath(root, q, &qp) {
         return nil
     }

     var i = len(pp) - 1
     var j = len(qp) - 1

     for i >= 0 && j >= 0 && pp[i] == qp[j] {
         i--
         j--
     }

     return pp[i + 1]
}


func getPath(root *TreeNode, p *TreeNode, path *[]*TreeNode) bool {
    if root == nil {
        return false
    }

    if root == p || getPath(root.Left, p, path) || getPath(root.Right, p, path) {
        *path = append(*path, root)
        return true
    }

    return false
}