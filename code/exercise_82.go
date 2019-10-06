/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil  || head.Next == nil{
        return head
    }
    // check head
    isMove := false
    for head.Next!=nil {
        if head.Val == head.Next.Val {
            isMove = true
        } else {
            if !isMove {
                break
            }
            isMove = false
        }
        head = head.Next
    }
    if isMove {
        head = head.Next
    }
    
    if head == nil  || head.Next == nil{
        return head
    }
   
    // check body
    last := head
    spot := last.Next
    flag := true
    for spot.Next != nil {
        if spot.Val != spot.Next.Val{
            if flag {
                last.Next = spot
                last = last.Next
            }
            flag = true
        } else {
            flag = false
        }
        spot = spot.Next
    }
    if flag {
        last.Next = spot
    } else {
        last.Next = nil
    }
    return head
}
