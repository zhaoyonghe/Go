type ListNode struct {
    Val int
    Next * ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    var nh *ListNode = &ListNode{0, nil}
    var cur *ListNode = nh
    
    for l1 != nil && l2 != nil {
        if l1.Val > l2.Val {
            cur.Next = l2
            l2 = l2.Next
        } else {
            cur.Next = l1
            l1 = l1.Next
        }
        cur = cur.Next
    }
    
    if l1 == nil {
        cur.Next = l2
    } else {
        cur.Next = l1
    }
    
    return nh.Next
}