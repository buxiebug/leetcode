package cases


/**
给定一个二叉树，返回其按层次遍历的节点值。 （即逐层地，从左到右访问所有节点）。

例如:
给定二叉树: [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其层次遍历结果：

[
  [3],
  [9,20],
  [15,7]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-level-order-traversal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
	var q []*TreeNode
	var result [][]int
	if root == nil {return result}
	q = append(q, root)
	left, right := 0, 1
	for left < right {
		tmp := make([]int, right-left)
		rightNew := right
		start := left
		for left < right {
			tmp[left - start] = q[left].Val
			if q[left].Left != nil { q = append(q, q[left].Left); rightNew += 1}
			if q[left].Right != nil {q = append(q, q[left].Right); rightNew += 1}
			left += 1
		}
		right = rightNew
		result = append(result, tmp)
	}
	return result
}
