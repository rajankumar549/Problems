/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func k344(head *ListNode,curr *ListNode, n int)( int, *ListNode){
    if curr == nil {
        return 0,nil
    }
    height ,nexxt:= k344(head,curr.Next,n)
    if height == n {
        curr.Next = nexxt
    }
    return height +1, curr.Next
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    height ,nxt:= k344(head,head,n)
    if height == n {
        head = nxt
    }
    return head
}
