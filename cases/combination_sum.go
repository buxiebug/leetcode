package cases

import (
	"fmt"
	"sort"
)

/**
给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的数字可以无限制重复被选取。

说明：

所有数字（包括 target）都是正整数。
解集不能包含重复的组合。 
示例 1:

输入: candidates = [2,3,6,7], target = 7,
所求解集为:
[
  [7],
  [2,2,3]
]
示例 2:

输入: candidates = [2,3,5], target = 8,
所求解集为:
[
  [2,2,2,2],
  [2,3,3],
  [3,5]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/combination-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func combination(candidates []int, idx int, target int) [][]int {
	var result [][]int
	if idx == len(candidates) {
		return [][]int{}
	}
	if target < 0 {
		return [][]int{}
	}
	if candidates[idx] == target {
		return [][]int{[]int{target}}
	}

	p1 := candidates[idx]
	p2 := combination(candidates, idx, target-p1)

	for _, subItem := range p2 {
		tmp := append([]int{candidates[idx]}, subItem...)
		result = append(result, tmp)
	}

	p3 := combination(candidates, idx+1, target)
	for _, subItem := range p3 {
		result = append(result, subItem)
	}
	return result
}

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	return combination(candidates, 0, target)
}

func TestCombinationSum(){
	//candidates := []int{2, 3, 6, 7}
	//target := 7
	candidates := []int{2, 5, 8, 4}
	target := 10
	fmt.Println(combinationSum(candidates, target))
}
