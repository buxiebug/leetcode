package cases

/**
给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

示例:

输入: [-2,1,-3,4,-1,2,1,-5,4],
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
进阶:

如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的分治法求解。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/maximum-subarray
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
/**
https://leetcode-cn.com/problems/maximum-subarray/solution/zui-da-zi-xu-he-cshi-xian-si-chong-jie-fa-bao-li-f/
 */

func maxSubArray(nums []int) int {
	max := nums[0]
	sum := 0
	idx := 0
	for idx < len(nums){
		sum += nums[idx]
		if sum > max { max = sum }
		if sum < 0 { sum = 0}
		idx += 1
	}
	return max
}

func maxSubArrayDP(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	max := dp[0]
	idx := 1
	for idx < len(nums){
		dp[idx] = nums[idx]
		if nums[idx] + dp[idx-1] > nums[idx] {
			dp[idx] = dp[idx-1] + nums[idx]
		}

		if max < dp[idx] {
			max = dp[idx]
		}
		idx += 1
	}
	return max
}
