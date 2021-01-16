type ListNode struct {
    Val int
    Next * ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
    var n int = len(lists)
    
    if n == 0 {
        return nil
    }
    
    for n != 1 {
        var span int = n / 2
        for i := 0; i < span; i++ {
            lists[i] = mergeTwoLists(lists[i], lists[i + span])
        }
        if n % 2 == 1 {
            lists[span] = lists[n - 1]
            n = (n + 1) / 2
        } else {
            n /= 2
        }
    }
    
    return lists[0]
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