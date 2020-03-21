type ListNode struct {
    Val int
    Next * ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    var dummy *ListNode = &ListNode{0, head}
    
    var a *ListNode = dummy
    var b *ListNode = dummy
    
    
    for i := 0; i < n; i++ {
        a = a.Next
    }
    
    for a.Next != nil {
        a = a.Next
        b = b.Next
    }
    
    b.Next = b.Next.Next
    
    return dummy.Next
}