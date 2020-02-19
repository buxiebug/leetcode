package cases

import "fmt"

/**
给定一个非负整数数组，你最初位于数组的第一个位置。

数组中的每个元素代表你在该位置可以跳跃的最大长度。

判断你是否能够到达最后一个位置。

示例 1:

输入: [2,3,1,1,4]
输出: true
解释: 我们可以先跳 1 步，从位置 0 到达 位置 1, 然后再从位置 1 跳 3 步到达最后一个位置。
示例 2:

输入: [3,2,1,0,4]
输出: false
解释: 无论怎样，你总会到达索引为 3 的位置。但该位置的最大跳跃长度是 0 ， 所以你永远不可能到达最后一个位置。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/jump-game
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

/**
https://leetcode-cn.com/problems/jump-game/solution/55-by-ikaruga/
 */

func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	i := 1
	for i < len(nums)-1 {
		if dp[i-1] <=0 {
			return false
		}
		dp[i] = dp[i-1] - 1
		if nums[i] > dp[i] {
			dp[i] = nums[i]
		}
		if dp[i] <= 0 {
			return false
		}
		i += 1
	}
	if dp[0] > 0 {
		return true
	} else {
		return false
	}
}

func TestCanJump() {
	t := []int{1, 2, 3}
	fmt.Println(canJump(t))
}
