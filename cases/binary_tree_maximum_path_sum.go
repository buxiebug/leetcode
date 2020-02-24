package cases

import "fmt"

/**
给定一个非空二叉树，返回其最大路径和。

本题中，路径被定义为一条从树中任意节点出发，达到任意节点的序列。该路径至少包含一个节点，且不一定经过根节点。

示例 1:

输入: [1,2,3]

       1
      / \
     2   3

输出: 6
示例 2:

输入: [-10,9,20,null,null,15,7]

   -10
   / \
  9  20
    /  \
   15   7

输出: 42

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-maximum-path-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func maxPathSumRecursion(root *TreeNode, max *int) int {
	if root == nil { return 0 }
	left := 0
	right := 0
	leftTmp := maxPathSumRecursion(root.Left, max) 
	if leftTmp> left { left = leftTmp}
	rightTmp := maxPathSumRecursion(root.Right, max) 
	if rightTmp> right { right = rightTmp}
	if root.Val + left + right > *max {
		*max = root.Val + left + right 
	}
	biggerBranch := left
	if right > left { biggerBranch =  right}
	return root.Val + biggerBranch
}

func maxPathSum(root *TreeNode) int {
	max := -1 << 31
	maxPathSumRecursion(root, &max)
	return max
}

func TestMaxPathSum() {
	n1 := &TreeNode{Val: -10}
	n2 := &TreeNode{Val: 9}
	n3 := &TreeNode{Val: 20}
	n4 := &TreeNode{Val: 15}
	n5 := &TreeNode{Val: 7}
	n1.Left = n2
	n1.Right = n3
	n3.Left = n4
	n3.Right = n5
	fmt.Println(maxPathSum(n1))
}
