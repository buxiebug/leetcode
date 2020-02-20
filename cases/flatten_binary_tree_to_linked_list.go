package cases

/**
给定一个二叉树，原地将它展开为链表。

例如，给定二叉树

    1
   / \
  2   5
 / \   \
3   4   6
将其展开为：

1
 \
  2
   \
    3
     \
      4
       \
        5
         \
          6

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func flatten(root *TreeNode) {
	stack := Stack{container: []interface{}{}}
	cur := root
	var last *TreeNode
	for stack.Len() > 0 || cur != nil {
		for cur != nil {
			stack.Push(cur)
			last = cur
			cur = cur.Left
		}
		top, _ := stack.Pop()
		cur = top.(*TreeNode)
		tmp := cur.Right
		if last != nil && last != cur {
			last.Right = cur.Right
			cur.Right = cur.Left
			cur.Left = nil
		}
		cur = tmp
	}
}

func TestFlatten() {
	n1 := &TreeNode{Val: 1}
	n2 := &TreeNode{Val: 2}
	n3 := &TreeNode{Val: 3}
	n4 := &TreeNode{Val: 4}
	n5 := &TreeNode{Val: 5}
	n6 := &TreeNode{Val: 6}
	n1.Left = n2
	n1.Right = n5
	n2.Left = n3
	n2.Right = n4
	n5.Right = n6
	flatten(n1)
}
