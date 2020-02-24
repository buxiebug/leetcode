package cases

import "fmt"

/**
在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序。

示例 1:

输入: 4->2->1->3
输出: 1->2->3->4
示例 2:

输入: -1->5->3->4->0
输出: -1->0->3->4->5

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sort-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// 归并　递归　｜｜　自底向上 https://leetcode-cn.com/problems/sort-list/solution/sort-list-gui-bing-pai-xu-lian-biao-by-jyd/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 这个解法用了递归，不考虑栈空间的消耗

func mergeList(l1 *ListNode, l2 *ListNode) *ListNode {
	dummpy := &ListNode{}
	p := dummpy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			p.Next = l1
			l1 = l1.Next
		} else {
			p.Next = l2
			l2 = l2.Next
		}
		p = p.Next
	}
	if l1 == nil {
		p.Next = l2
	} else {
		p.Next = l1
	}
	return dummpy.Next
}

func findMid(l *ListNode) *ListNode {
	dummy := &ListNode{}
	dummy.Next = l
	fast, slow := dummy, dummy
	for fast != nil && fast.Next != nil{
		fast = fast.Next
		slow = slow.Next
		if fast != nil {
			fast = fast.Next
		}
	}
	return slow
}

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	mid := findMid(head)
	if mid == head{
		if mid.Next == nil {return head}
		next := mid.Next
		mid.Next = nil
		return mergeList(mid, next)
	}
	
	right := sortList(mid.Next)
	mid.Next = nil
	left := sortList(head)
	return mergeList(left, right)
}

func TestSortList() {
	n1 := &ListNode{Val: 4}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 1}
	n4 := &ListNode{Val: 3}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4

	n := sortList(n1)
	for n != nil {
		fmt.Println(n.Val)
		n = n.Next
	}
}
