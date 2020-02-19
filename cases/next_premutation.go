package cases

import "fmt"

/**
实现获取下一个排列的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。

如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。

必须原地修改，只允许使用额外常数空间。

以下是一些例子，输入位于左侧列，其相应输出位于右侧列。
1,2,3 → 1,3,2
3,2,1 → 1,2,3
1,1,5 → 1,5,1

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/next-permutation
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func nextPermutation(nums []int)  {
	if len(nums) <2 {return}
	i := len(nums) - 1
	MinGtIdx := i
	for i > 0 {
		j := i-1
		if nums[j] < nums[i] {
			k := len(nums) - 1
			MinGtIdx = i
			for k > j {
				if nums[k] > nums[j] && nums[MinGtIdx] > nums[k] { MinGtIdx = k}
				k -= 1
			}
			nums[MinGtIdx], nums[j] = nums[j], nums[MinGtIdx]
			m := j+1
			for m < (len(nums) -1) {  // case : [2,3,1,3,3] => [2,3,3,1,3]
				if nums[m] < nums[m+1] {
					nums[m], nums[m+1] = nums[m+1], nums[m]
				} else {break}
				m += 1
			}
			m = 0
			for m < (len(nums) -1 -j)/2{
				nums[j+1+m], nums[len(nums)-1-m ] = nums[len(nums)-1-m], nums[j+1+m]
				m += 1
			}
			break
		}
		i -= 1
	}
	if i == 0 {
		m := 0
		for m < len(nums) / 2 {
			nums[m], nums[len(nums)-1-m ] = nums[len(nums)-1-m], nums[m]
			m += 1
		}
	}
}


func TestNextPremutation() {
	t := []int{2, 3, 1, 3, 3 }
	nextPermutation(t)
	fmt.Println(t)
}
