package cases

import (
	"fmt"
	"sort"
)

/**
给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。

注意：答案中不可以包含重复的三元组。

示例：

给定数组 nums = [-1, 0, 1, 2, -1, -4]，

满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/3sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

/*
solution:
https://leetcode-cn.com/problems/3sum/solution/pai-xu-shuang-zhi-zhen-zhu-xing-jie-shi-python3-by/
 */

func threeSum(nums []int) [][]int {
	var rs [][]int
	rs = make([][]int, 0)
	sort.Ints(nums)
	for i:=0; i<len(nums); i++{
		if nums[i] > 0 { return rs }
		if i > 0 && nums[i] == nums[i-1]{ continue }
		j, k := i+1, len(nums)-1
		for j < k {
			if nums[i] + nums[j] + nums[k] == 0{
				rs = append(rs, []int{nums[i], nums[j], nums[k]})
				for j<k && nums[j] == nums[j+1]{
					j += 1
				}
				for j<k && nums[k] == nums[k-1]{
					k -= 1
				}
				j += 1
				k -= 1
			} else if nums[i] + nums[j] + nums[k] < 0{
				j += 1
			} else {
				k -= 1
			}
		}
	}
	return rs
}

func TestThreeSum(){
	sums := []int{0, 0, 0}
	fmt.Println(threeSum(sums))
}
