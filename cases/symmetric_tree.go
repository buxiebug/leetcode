package cases

/**
给定一个二叉树，检查它是否是镜像对称的。

例如，二叉树 [1,2,2,3,4,4,3] 是对称的。

    1
   / \
  2   2
 / \ / \
3  4 4  3
但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:

    1
   / \
  2   2
   \   \
   3    3
说明:

如果你可以运用递归和迭代两种方法解决这个问题，会很加分。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/symmetric-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

/**
迭代是逐渐逼近，用新值覆盖旧值，直到满足条件后结束，不保存中间值，空间利用率高。
递归是将一个问题分解为若干相对小一点的问题，遇到递归出口再原路返回，因此必须保存相关的中间值，这些中间值压入栈保存，问题规模较大时会占用大量内存。
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetricNoRecursive(root *TreeNode) bool {
	s1 := Stack{container: []interface{}{}}
	s2 := Stack{container: []interface{}{}}
	cur1 := root
	cur2 := root

	for (s1.Len() > 0 && s2.Len() > 0) || (cur1 != nil && cur2 != nil) {
		for cur1 != nil {
			s1.Push(cur1)
			cur1 = cur1.Left
		}
		for cur2 != nil {
			s2.Push(cur2)
			cur2 = cur2.Right
		}
		if s1.Len() != s2.Len() { return false }
		top1, _ := s1.Pop()
		top2, _ := s2.Pop()
		cur1 = top1.(*TreeNode)
		cur2 = top2.(*TreeNode)
		if cur1.Val != cur2.Val { return false }
		cur1 = cur1.Right
		cur2 = cur2.Left
	}
	if s1.Len() != s2.Len() {return false}
	if cur1 != cur2 {return false}
	return true
}

func isMirror(node1 *TreeNode, node2 *TreeNode) bool{
	if node1 == nil || node2 == nil {
		if node1 == node2 {return true}
		return false
	}
	if node1.Val != node2.Val { return false}
	return isMirror(node1.Left, node2.Right) && isMirror(node1.Right, node2.Left)
}

func isSymmetricRecursive(root *TreeNode) bool {
	if root == nil {return true}
	return isMirror(root.Left, root.Right)
}
