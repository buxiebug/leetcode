package cases

/**
设计一个支持 push，pop，top 操作，并能在常数时间内检索到最小元素的栈。

push(x) -- 将元素 x 推入栈中。
pop() -- 删除栈顶的元素。
top() -- 获取栈顶元素。
getMin() -- 检索栈中的最小元素。
示例:

MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin();   --> 返回 -3.
minStack.pop();
minStack.top();      --> 返回 0.
minStack.getMin();   --> 返回 -2.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/min-stack
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type StackNode struct {
	Min int
	Val int
}

type MinStack struct {
	min   int
	nodes []StackNode
}

/** initialize your data structure here. */
func ConstructorMinStack() MinStack {
	return MinStack{
		min:   1<<31 - 1,
		nodes: []StackNode{},
	}
}

func (this *MinStack) Push(x int) {
	if this.min > x {
		this.min = x
	}
	this.nodes = append(this.nodes, StackNode{
		Min: this.min,
		Val: x,
	})
}

func (this *MinStack) Pop() {
	if len(this.nodes) == 0 {return}
	this.nodes = this.nodes[:len(this.nodes)-1]
	if len(this.nodes) > 0 {
		this.min = this.nodes[len(this.nodes)-1].Min
	} else {
		this.min = 1 << 31 - 1
	}
}

func (this *MinStack) Top() int {
	if len(this.nodes) == 0 {return 0}
	top := this.nodes[len(this.nodes)-1]
	return top.Val
}

func (this *MinStack) GetMin() int {
	if len(this.nodes) == 0 {return 0}
	top := this.nodes[len(this.nodes)-1]
	return top.Min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
