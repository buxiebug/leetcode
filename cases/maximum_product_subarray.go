package cases

/**
给定一个整数数组 nums ，找出一个序列中乘积最大的连续子序列（该序列至少包含一个数）。

示例 1:

输入: [2,3,-2,4]
输出: 6
解释: 子数组 [2,3] 有最大乘积 6。
示例 2:

输入: [-2,0,-1]
输出: 0
解释: 结果不能为 2, 因为 [-2,-1] 不是子数组。
*/
//

func maxInt(a int, b int) int{
	if a > b { return a }
	return b
}

func minInt(a int, b int) int{
	if a > b { return b }
	return a
}

func maxProduct(nums []int) int {
	max := nums[0]
	curMax := nums[0]
	curMin := nums[0]
	for i := 1; i < len(nums); i++ {
		curMaxTmp := maxInt(maxInt(nums[i], curMax * nums[i]), curMin * nums[i])
		curMin = minInt(minInt(nums[i], curMin * nums[i]), curMax * nums[i])
		curMax = curMaxTmp
		if max < curMax {
			max = curMax
		}
	}
	return max
}
