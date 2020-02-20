package cases

/**
根据一棵树的前序遍历与中序遍历构造二叉树。

注意:
你可以假设树中没有重复的元素。

例如，给出

前序遍历 preorder = [3,9,20,15,7]
中序遍历 inorder = [9,3,15,20,7]
返回如下的二叉树：

    3
   / \
  9  20
    /  \
   15   7

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal
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
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 { return nil }
	node := &TreeNode{Val: preorder[0]}
	idx := 0
	for idx < len(inorder) {
		if inorder[idx] == preorder[0] {break}
		idx += 1
	}
	if len(preorder) > 1 {
		node.Left = buildTree(preorder[1:idx+1], inorder[:idx])
	} else {
		node.Left = nil
	}
	if idx < len(preorder) - 1 {
		node.Right = buildTree(preorder[idx+1:], inorder[idx+1:])
	} else {
		node.Right = nil
	}
	return node
}
