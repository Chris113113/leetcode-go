package main

import ("log")

type ListNode struct {
    Val int
    Next *ListNode
}

func main() {
    log.Println(addTwoNumbers(&ListNode{2, &ListNode{4, &ListNode{3, nil}}}, &ListNode{5, &ListNode{6, &ListNode{4, nil}}}))
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    return recursiveAddTwoNumbers(l1, l2, 0)
}

func recursiveAddTwoNumbers(l1 *ListNode, l2 *ListNode, carry int) *ListNode {
    if l1 == nil && l2 == nil {
        if carry > 0 {
            return &ListNode{carry, nil}
        }
        
        return nil
    }
    
    if l1 == nil {
        l1 = &ListNode{0, nil}
    }
    
    if l2 == nil {
        l2 = &ListNode{0, nil}
    }
    
    sum := l1.Val + l2.Val + carry
    log.Println(sum, sum%10, sum/10)
    return &ListNode{sum % 10, recursiveAddTwoNumbers(l1.Next, l2.Next, sum / 10 )}
}