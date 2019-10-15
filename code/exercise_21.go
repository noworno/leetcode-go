/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil && l2 != nil {
        return l2
    } else if l1 != nil && l2 == nil {
        return l1
    } else if l1 == nil && l2 == nil {
        return nil
    }
    if l1.Val > l2.Val {
        l1,l2=l2,l1
    }
    idx1 := l1
    idx2 := l2
    for idx2 != nil {
        if idx1.Next == nil {
            idx1.Next = idx2
            break
        }
        if idx1.Val<=idx2.Val && idx1.Next.Val > idx2.Val {
            t := idx1.Next
            idx1.Next = idx2
            idx2 = idx2.Next
            idx1.Next.Next = t
            idx1 = idx1.Next
        } else {
            idx1 = idx1.Next
        }
    }
    return l1
}