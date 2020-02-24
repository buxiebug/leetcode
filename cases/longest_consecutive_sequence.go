package cases

/**
给定一个未排序的整数数组，找出最长连续序列的长度。

要求算法的时间复杂度为 O(n)。

示例:

输入: [100, 4, 200, 1, 3, 2]
输出: 4
解释: 最长连续序列是 [1, 2, 3, 4]。它的长度为 4。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-consecutive-sequence
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/**
https://leetcode-cn.com/problems/longest-consecutive-sequence/solution/dong-tai-gui-hua-python-ti-jie-by-jalan/
https://leetcode-cn.com/problems/longest-consecutive-sequence/solution/zui-chang-lian-xu-xu-lie-by-leetcode/
*/


func longestConsecutive(nums []int) int {
	m := map[int]int{}
	max := 0
	for _, num := range nums {
		_, ok := m[num]
		if ok { continue }
		left, leftOk := m[num-1]
		if !leftOk { left=0 }
		right, rightOk := m[num+1]
		if !rightOk { right=0 }
		length := 1 + left + right
		if max < length { max = length}
		m[num] = length
		m[num-left] = length
		m[num+right] = length
	}
	return max
}