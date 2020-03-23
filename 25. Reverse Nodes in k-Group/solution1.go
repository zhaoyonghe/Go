type ListNode struct {
    Val int
    Next * ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
    if head == nil {
        return head
    }
    
    var dummy *ListNode = &ListNode{0, head}
    var cur *ListNode = dummy
    
    var stub *ListNode = &ListNode{0, nil}
    var top *ListNode = stub
    var a *ListNode = head
    var b *ListNode = head.Next
    
    for {
        var i int = 0
        var count *ListNode = cur
        
        for count.Next != nil && i < k {
            count = count.Next
            i++
        }
        
        if i < k {
            return dummy.Next
        } else {
            i = 0
        }
        
        var mark *ListNode = a
        
        for i < k {
            a.Next = top
            top = a

            a = b
            if b == nil {
                break
            }
            b = b.Next 
            
            i++
        }
        
        cur.Next = top
        mark.Next = a
        cur = mark
    }
    
    return dummy.Next
}