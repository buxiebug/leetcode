package cases

import "fmt"

/**
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-two-numbers
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

type ListNode struct {
    Val int
    Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	step := 0
	iter1 := l1
	iter2 := l2
	var head *ListNode
	var cur *ListNode

	for{
		if iter1 == nil && iter2 == nil {break}
		iter1Value := 0
		iter2Value := 0
		if iter1 != nil {iter1Value = iter1.Val}
		if iter2 != nil {iter2Value = iter2.Val}

		s := iter1Value + iter2Value + step
		step = int(s / 10)
		num := s % 10
		if head == nil {
			head = &ListNode{
				Val: num,
			}
			cur = head
		} else {
			cur.Next = &ListNode{
				Val: num,
			}
			cur = cur.Next
		}
		if iter1 !=nil {
			iter1 = iter1.Next
		}
		if iter2 != nil {
			iter2 = iter2.Next
		}
	}
	if step > 0 {
		cur.Next = &ListNode{
			Val:  step,
			Next: nil,
		}
	}
	return head
}

func TestTwoSum1() {
	n11 := &ListNode{Val:2}
	n12 := &ListNode{Val:4}
	n13 := &ListNode{Val:3}
	n21 := &ListNode{Val:5}
	n22 := &ListNode{Val:6}
	n23 := &ListNode{Val:4}
	n24 := &ListNode{Val:1}
	n11.Next = n12
	n12.Next = n13
	n21.Next = n22
	n22.Next = n23
	n23.Next = n24
	rs := addTwoNumbers(n11, n21)
	for{
		if rs == nil{break}
		fmt.Println(rs.Val)
		rs = rs.Next
	}
}
