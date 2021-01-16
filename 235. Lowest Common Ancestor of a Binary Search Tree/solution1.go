type TreeNode struct {
    Val int
    Left *ListNode
    Right *ListNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    if p.Val > q.Val {
        return lowestCommonAncestor(root, q, p)
    }

    var cur *TreeNode = root

    for {
        if p.Val <= cur.Val && cur.Val <= q.Val {
            return cur
        }
        if cur.Val < p.Val {
            cur = cur.Right
        }
        if cur.Val > q.Val {
            cur = cur.Left
        }
    }

    return nil
}