package cases

import "fmt"

/**
给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。

求在该柱状图中，能够勾勒出来的矩形的最大面积。

以上是柱状图的示例，其中每个柱子的宽度为 1，给定的高度为 [2,1,5,6,2,3]。

图中阴影部分为所能勾勒出的最大矩形面积，其面积为 10 个单位。

示例:

输入: [2,1,5,6,2,3]
输出: 10

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/largest-rectangle-in-histogram
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

/**
看了别人的答案想了半天才明白……其实可以把这个想象成锯木板，如果木板都是递增的那我很开心，如果突然遇到一块木板i矮了一截，那我就先找之前最戳出来的一块（其实就是第i-1块），计算一下这个木板单独的面积，然后把它锯成次高的，这是因为我之后的计算都再也用不着这块木板本身的高度了。再然后如果发觉次高的仍然比现在这个i木板高，那我继续单独计算这个次高木板的面积（应该是第i-1和i-2块），再把它俩锯短。直到发觉不需要锯就比第i块矮了，那我继续开开心心往右找更高的。当然为了避免到了最后一直都是递增的，所以可以在最后加一块高度为0的木板。
单调栈
https://leetcode-cn.com/problems/largest-rectangle-in-histogram/solution/084zhu-zhuang-tu-zhong-zui-da-de-ju-xing-by-6westb/
 */


func largestRectangleArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}

	heights = append(heights, 0)
	heights = append([]int{0}, heights...)

	max := 0
	stack := Stack{container: []interface{}{}}
	for idx:=0; idx < len(heights); idx +=1 {
		for stack.Len() > 0 {
			top, _:= stack.Pop()
			if heights[top.(int)] > heights[idx]{
				tmp := (idx - top.(int) - 1) * heights[top.(int)]
				if tmp > max { max = tmp }
			} else {
				stack.Push(top)
				break
			}
		}
		stack.Push(idx)

	}
	return max
}

func TestLargestRectangleInHistogram() {
	t := []int{2,1,5,6,2,3}
	fmt.Println(largestRectangleArea(t))
}