type ListNode struct {
    Val int
    Next * ListNode
}

func addTwoNumbers(l1 * ListNode, l2 * ListNode) * ListNode {
    var dummy * ListNode = & ListNode {
        0, nil
    }
    var cur * ListNode = dummy

    var sum int = 0
    var carry int = 0

    for l1 != nil || l2 != nil {
        var l1val int = 0
        var l2val int = 0

        if l1 != nil {
            l1val = l1.Val
            l1 = l1.Next
        }

        if l2 != nil {
            l2val = l2.Val
            l2 = l2.Next
        }

        var temp int = (l1val + l2val + carry)
        sum = temp % 10
        carry = temp / 10

        cur.Next = & ListNode {
            sum, nil
        }
        cur = cur.Next

        //fmt.Println(dummy.Next)
    }

    if carry == 1 {
        cur.Next = & ListNode {
            1, nil
        }
    }

    return dummy.Next
}