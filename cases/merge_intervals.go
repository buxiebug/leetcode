package cases

import "sort"

/**
给出一个区间的集合，请合并所有重叠的区间。

示例 1:

输入: [[1,3],[2,6],[8,10],[15,18]]
输出: [[1,6],[8,10],[15,18]]
解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
示例 2:

输入: [[1,4],[4,5]]
输出: [[1,5]]
解释: 区间 [1,4] 和 [4,5] 可被视为重叠区间。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/merge-intervals
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool{
		if intervals[i][0] < intervals[j][0] {
			return true
		}
		return false
	})
	var rs [][]int
	if len(intervals) <= 1 { return intervals }
	idx :=  1
	tmpLeft, tmpRight := intervals[0][0], intervals[0][1]
	for idx < len(intervals) {
		interval := intervals[idx]
		idx += 1

		if tmpRight < interval[0] {
			rs = append(rs, []int{tmpLeft, tmpRight})
			tmpLeft = interval[0]
			tmpRight = interval[1]
		} else {
			if interval[0] < tmpLeft {
				tmpLeft = interval[0]
			}
			if interval[1] > tmpRight {
				tmpRight = interval[1]
			}
		}
	}
	rs = append(rs, []int{tmpLeft, tmpRight})
	return rs
}
