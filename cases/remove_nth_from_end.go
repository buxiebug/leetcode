package cases

import "fmt"

/**
给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。

示例：

给定一个链表: 1->2->3->4->5, 和 n = 2.

当删除了倒数第二个节点后，链表变为 1->2->3->5.
说明：

给定的 n 保证是有效的。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	step := 0
	fast := head
	slow := head
	pre := head
	for step < n {
		fast = fast.Next
		step += 1
	}
	for fast != nil {
		fast = fast.Next
		pre = slow
		slow = slow.Next
	}
	if pre == slow {
		head = head.Next
	} else {
		pre.Next = slow.Next
	}
	return head
}

func TestRemoveNthFromEnd() {
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 3}
	n4 := &ListNode{Val: 4}
	n5 := &ListNode{Val: 5}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5
	head := removeNthFromEnd(n1, 2)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
