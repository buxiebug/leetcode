package cases

import "fmt"

/**
给定一个二叉树，返回它的中序 遍历。

示例:

输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [1,3,2]
进阶: 递归算法很简单，你可以通过迭代算法完成吗

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-inorder-traversal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var rs []int
	stack := Stack{container:[]interface{}{}}
	cur := root
	for stack.Len() > 0 || cur != nil{
		if cur != nil {
			stack.Push(cur)
			cur = cur.Left
		} else {
			top, _ := stack.Pop()
			cur = top.(*TreeNode)
			rs = append(rs, cur.Val)
			cur = cur.Right
		}
	}
	return rs
}

func TestBinaryTreeInorderTraversal () {
	n1 := &TreeNode{Val:1}
	n2 := &TreeNode{Val:2}
	n3 := &TreeNode{Val:3}
	n1.Right = n2
	n2.Left = n3
	fmt.Println(inorderTraversal(n1))

	/**
	坑啊
	 */
	i := 1
	for j:= 0; j < 5; j++ {
		i := i + j
		fmt.Println(i)
	}
}
