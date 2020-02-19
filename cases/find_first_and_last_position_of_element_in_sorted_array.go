package cases

import "fmt"

/**
给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。

你的算法时间复杂度必须是 O(log n) 级别。

如果数组中不存在目标值，返回 [-1, -1]。

示例 1:

输入: nums = [5,7,7,8,8,10], target = 8
输出: [3,4]
示例 2:

输入: nums = [5,7,7,8,8,10], target = 6
输出: [-1,-1]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func searchPart(nums []int, left, right, target int) []int {
	rl, rr := -1, -1
	if left > right {
		return []int{rl, rr}
	}
	mid := (left + right) / 2
	if nums[mid] < target {
		return searchPart(nums, mid+1, right, target)
	}  else if nums[mid] > target {
		return searchPart(nums, left, mid-1, target)
	}
	lrs := searchPart(nums, left, mid-1, target)
	rrs := searchPart(nums, mid+1, right, target)
	rl, rr = mid, mid
	if lrs[0] != -1 { rl = lrs[0] }
	if rrs[1] != -1 { rr = rrs[1]}
	return []int{rl, rr}
}

func searchRange(nums []int, target int) []int {
	return searchPart(nums, 0, len(nums)-1, target)
}

func TestSearchRange(){
	nums := []int{5, 7, 7, 8, 8, 10}
	fmt.Println(searchRange(nums, 8))
}