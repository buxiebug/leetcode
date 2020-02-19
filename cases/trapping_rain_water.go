package cases

/**
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。 感谢 Marcos 贡献此图。

示例:

输入: [0,1,0,2,1,0,1,3,2,1,2,1]
输出: 6

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/trapping-rain-water
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
/**
https://leetcode-cn.com/problems/trapping-rain-water/solution/xiang-xi-tong-su-de-si-lu-fen-xi-duo-jie-fa-by-w-8/
 */

func trap(height []int) int {
	left, maxLeft, right, maxRight := 1, 0, len(height)-2, 0
	sum, cur := 0, 1
	for cur < len(height) - 1 {
		if height[left-1] > height[right+1] {
			if height[right+1] > maxRight { maxRight = height[right+1]}
			if maxRight > height[right] {
				sum += maxRight - height[right]
			}
			right -= 1
		} else {
			if height[left-1] > maxLeft { maxLeft = height[left-1]}
			if maxLeft > height[left] {
				sum += maxLeft - height[left]
			}
			left += 1
		}
		cur += 1
	}
	return sum
}
