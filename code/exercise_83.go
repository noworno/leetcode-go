/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil || head.Next == nil{
        return head
    }
    var l *ListNode = head
    for l.Next != nil {
        if l.Val == l.Next.Val {
             l.Next = l.Next.Next
        } else {
             l = l.Next
        }
    }
    return head
}
