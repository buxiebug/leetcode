package cases

import (
	"fmt"
)

/**
在未排序的数组中找到第 k 个最大的元素。请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。

示例 1:

输入: [3,2,1,5,6,4] 和 k = 2
输出: 5
示例 2:

输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
输出: 4
说明:

你可以假设 k 总是有效的，且 1 ≤ k ≤ 数组的长度。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/kth-largest-element-in-an-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func findKthLargest(nums []int, k int) int {
	if len(nums) == 2 {
		if nums[0] > nums[1] {
			nums[0], nums[1] = nums[1], nums[0]
		}
		return nums[2-k]
	}
	i, j := 1, len(nums)-1
	for {
		for i <= j && nums[i] <= nums[0] {
			i++
		}
		for i < j && nums[j] > nums[0] {
			j--
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		} else {
			break
		}
	}
	i--
	if k == len(nums)-i {
		return nums[0]
	}
	nums[0], nums[i] = nums[i], nums[0]
	if k < len(nums)-i {
		return findKthLargest(nums[i+1:], k)
	}
	return findKthLargest(nums[:i], k-len(nums)+i)
}

func TestFindKthLargest() {
	fmt.Println(findKthLargest([]int{3, 3, 3, 3, 3, 3, 3, 3, 3}, 8))
}
