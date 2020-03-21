type ListNode struct {
    Val int
    Next * ListNode
}

func swapPairs(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    
    var dummy *ListNode = &ListNode{0, head}
    
    var back *ListNode = dummy
    var a *ListNode = head
    var b *ListNode = head.Next
    
    
    for {
        a.Next = a.Next.Next
        b.Next = a
        back.Next = b
        
        back = a
        a = a.Next
        
        if a == nil || a.Next == nil {
            return dummy.Next
        }
        
        b = a.Next
    }
    
    return nil
}