package cases

/**
请判断一个链表是否为回文链表。

示例 1:

输入: 1->2
输出: false
示例 2:

输入: 1->2->2->1
输出: true
进阶：
你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/palindrome-linked-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	slow, fast := head, head
	left, right := head, head
	pre, next := new(ListNode), head
	for fast != nil {
		next = slow.Next
		if fast.Next == nil {
			left = pre
			right = next
			break
		}
		fast = fast.Next
		if fast.Next == nil {
			slow.Next = pre
			left = slow
			right = next
			break
		}
		fast = fast.Next

		slow.Next = pre
		pre = slow
		slow = next
	}
	for left != nil && right != nil {
		if right.Val != left.Val {return false}
		left = left.Next
		right = right.Next
	}
	return true
}